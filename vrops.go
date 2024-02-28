package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pdxfixit/hostdb"
)

// using an adapter ID, get the related resources
func getAdapterResources(adapterID string, currentPage int) (resources vropsAdapterResources, err error) {

	if err := resources.LoadFrom(
		fmt.Sprintf(
			"%s/suite-api/api/adapters/%s/resources?compression=enabled&page=%d&amp;pageSize=%d",
			config.Vrops.Host,
			adapterID,
			currentPage,
			config.Vrops.PageSize,
		),
	); err != nil {
		return vropsAdapterResources{}, err
	}

	return resources, nil

}

// get the properties for a slice of resources, return a slice of HostDB records
func getResourceProperties(resources []vropsResource) (collection []hostdb.Record) {

	// for each of the resources
	for i, resource := range resources {

		if config.Collector.Debug {
			log.Println(fmt.Sprintf(
				"Resource %d/%d (%s)...",
				i+1,
				len(resources),
				resource.ResourceKey.ResourceKindKey,
			))
			log.Println(fmt.Sprintf("%v", resource))
		}

		// if this is not a resource listed in the config, skip it
		wanted := false
		for _, kind := range config.Vrops.ResourceKindKeys {
			if kind == resource.ResourceKey.ResourceKindKey {
				wanted = true
				break
			}
		}

		if !wanted {
			log.Println("Skipping!")
			continue
		}

		vropsResourceProperties := vropsResourceProperties{}

		// get the properties for this resource
		if err := vropsResourceProperties.LoadFrom(
			fmt.Sprintf(
				"%s/suite-api/api/resources/%s/properties?compression=enabled",
				config.Vrops.Host,
				resource.Identifier,
			),
		); err != nil {
			log.Println(err)
			continue
		}

		if config.Collector.Debug {
			log.Println(fmt.Sprintf("Found %d properties for the resource %s.", len(vropsResourceProperties.Property), resource.Identifier))
		}

		// TODO: validate data

		// marshal into json
		jsonPayload, err := json.Marshal(vropsResourceProperties)
		if err != nil {
			log.Println(err)
			continue
		}

		// set the record type e.g. vrops-vmware-virtualmachine
		recordType := strings.ToLower(fmt.Sprintf("vrops-%s-%s",
			strings.Replace(resource.ResourceKey.AdapterKindKey, " ", "_", -1),
			strings.Replace(resource.ResourceKey.ResourceKindKey, " ", "_", -1),
		))

		// special type handling
		hostname := ""
		ip := ""
		switch recordType {
		case "vrops-vmware-hostsystem":
			for _, property := range vropsResourceProperties.Property {
				switch property.Name {
				case "config|name":
					hostname = property.Value
				case "net:vmk0|ip_address":
					ip = property.Value
				}
			}
		case "vrops-vmware-virtualmachine":
			for _, property := range vropsResourceProperties.Property {
				switch property.Name {
				case "summary|guest|hostName":
					if property.Value == "localhost" {
						continue
					}
					hostname = property.Value
				case "summary|guest|ipAddress":
					if property.Value == "127.0.0.1" {
						continue
					}
					ip = property.Value

				}
			}
		}

		// stow the whole thing in hostdb
		record := hostdb.Record{
			ID:        "",
			Type:      recordType,
			Hostname:  hostname,
			IP:        ip,
			Timestamp: time.Now().UTC().Format("2006-01-02 15:04:05"),
			Committer: "hostdb-collector-vrops",
			Context:   nil,
			Data:      jsonPayload,
			Hash:      "",
		}

		collection = append(collection, record)

	}

	return

}

// get a session token from vrops
func getSessionToken() string {

	log.Println(fmt.Sprintf("Trying %s...", config.Vrops.Host))
	session, err := httpRequest(
		"POST",
		fmt.Sprintf("%s/suite-api/api/auth/token/acquire", config.Vrops.Host),
		strings.NewReader(fmt.Sprintf("{\"username\":\"%s@pdxfixit.com\",\"password\":\"%s\"}", config.Vrops.User, config.Vrops.Pass)),
		vropsSessionHeaders,
	)
	if err != nil {
		log.Println(fmt.Sprintf("%s", session))
		log.Fatal(err)
	}

	// unmarshal the response into a struct
	vropsSession := vropsSessionToken{}
	if err := json.Unmarshal(session, &vropsSession); err != nil {
		log.Fatal(err)
	}

	if config.Collector.Debug {
		log.Println(fmt.Sprintf("vRealizeOpsToken %s", vropsSession.Token))
	}

	return vropsSession.Token

}

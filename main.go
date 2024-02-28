package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var vropsSessionHeaders = map[string]string{
	"Accept":       "application/json",
	"Content-Type": "application/json",
}

func main() {

	// load config
	loadConfig()

	// get a session token
	vropsSessionHeaders["Authorization"] = fmt.Sprintf(
		"vRealizeOpsToken %s",
		getSessionToken(),
	)

	log.Println(fmt.Sprintf(
		"Getting a list of vCenters from %s...",
		config.Vrops.Host,
	))

	vropsAdapterList := vropsAdapterList{}

	// collect a list of vcenter instances from vrops
	if err := vropsAdapterList.LoadFrom(
		fmt.Sprintf(
			"%s/suite-api/api/adapters?compression=enabled",
			config.Vrops.Host,
		),
	); err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf(
		"Found %d adapter instances. Not all of these are vcenters.",
		len(vropsAdapterList.Instances),
	))

	// for each of the adapter instances
	for n, adapter := range vropsAdapterList.Instances {

		// if it's not a vcenter, move on to the next
		if adapter.ResourceKey.AdapterKindKey != "VMWARE" {
			continue
		}

		log.Println(fmt.Sprintf(
			"Adapter %d/%d (%s)...",
			n+1,
			len(vropsAdapterList.Instances),
			adapter.ResourceKey.Name,
		))
		if config.Collector.Debug {
			log.Println(fmt.Sprintf("%v", adapter))
		}

		vropsAdapterResources, err := getAdapterResources(adapter.ID, 0)
		if err != nil {
			log.Println(err)
			continue
		}

		// check the number of resources
		log.Println(fmt.Sprintf(
			"Found %d resources for the adapter %s. Only the ResourceKindKeys listed in config will be collected.",
			vropsAdapterResources.PageInfo.TotalCount,
			adapter.ID,
		))

		// collect the first page of resources
		records := getResourceProperties(vropsAdapterResources.ResourceList)

		// figure out how many iterations we need total
		iterations := vropsAdapterResources.PageInfo.TotalCount / config.Vrops.PageSize
		if (vropsAdapterResources.PageInfo.TotalCount % config.Vrops.PageSize) > 0 {
			iterations++
		}

		// if >1k resources then we gotta navigate some pagination
		for i := 1; i < iterations; i++ {

			log.Println(fmt.Sprintf(
				"Iteration %d/%d...",
				i+1,
				iterations,
			))

			resources, err := getAdapterResources(adapter.ID, i)
			if err != nil {
				log.Println(err)
				continue
			}

			// collect the resources
			records = append(records, getResourceProperties(resources.ResourceList)...)

		}

		// create a recordset
		recordSet := createRecordSet(adapter, records)

		// post to HostDB
		if config.Collector.SampleData {
			if err := recordSet.Save(fmt.Sprintf("/sample-data/%s.json", recordSet.Context["vc_url"])); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := recordSet.Send(fmt.Sprintf("vc_url=%s", recordSet.Context["vc_url"])); err != nil {
				log.Fatal(err)
			}
		}

	}

	// TODO: Waiting on TVE-366
	// logout from vrops, destroy session
	//log.Println("Closing vrops session...")
	//if deleteResponse, err := httpRequest(
	//	"POST",
	//	fmt.Sprintf("%s/suite-api/api/auth/token/release", config.Vrops.Host),
	//	nil,
	//	vropsSessionHeaders,
	//); err != nil {
	//	log.Println(fmt.Sprintf("%s", deleteResponse))
	//	log.Println(err) // don't die just because we couldn't expire the token
	//}

	log.Println("All done!")

}

func httpRequest(method string, url string, body io.Reader, header map[string]string) (bytes []byte, err error) {

	var res *http.Response

	// Client
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	// INSECURE
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// headers
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	res, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		return bytes, errors.New(res.Status)
	}

	return bytes, nil

}

func requestToStruct(url string, obj interface{}) (err error) {

	response, err := httpRequest(
		"GET",
		url,
		nil,
		vropsSessionHeaders,
	)
	if err != nil {
		log.Println(fmt.Sprintf("%v", response))
		return err
	}

	// unmarshal the response into a struct
	if err := json.Unmarshal(response, &obj); err != nil {
		log.Fatal(err)
		return err // i know this is kinda pointless, but i want it here
	}

	return nil

}

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pdxfixit/hostdb"
)

func createRecordSet(adapter vropsAdapterInstance, records []hostdb.Record) (recordSet hostdb.RecordSet) {

	// context
	context := map[string]interface{}{
		"vc_name": adapter.ResourceKey.Name,
	}

	// attempt to get a vCenter URL
	for _, identifier := range adapter.ResourceKey.ResourceIdentifiers {
		if identifier.IdentifierType.Name == "VCURL" {
			context["vc_url"] = identifier.Value
			break
		}
	}

	// if there's a description
	if adapter.Description != "" {
		context["vc_desc"] = adapter.Description
	}

	recordSet = hostdb.RecordSet{
		Type: strings.ToLower(fmt.Sprintf(
			"vrops-%s",
			strings.Replace(adapter.ResourceKey.AdapterKindKey, " ", "_", -1),
		)), // e.g. vrops-vmware
		Timestamp: time.Now().UTC().Format("2006-01-02 15:04:05"),
		Context:   context,
		Committer: "hostdb-collector-vrops",
		Records:   records,
	}

	return recordSet

}

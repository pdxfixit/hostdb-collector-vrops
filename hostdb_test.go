package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pdxfixit/hostdb"
	"github.com/stretchr/testify/assert"
)

// given a vropsAdapterInstance, and a slice of hostdb.Records, return a hostdb.RecordSet
func TestCreateRecordSet(t *testing.T) {

	adapter := vropsAdapterInstance{
		ResourceKey: vropsResourceKey{
			Name:            "test",
			AdapterKindKey:  "VMWARE",
			ResourceKindKey: "VMwareAdapter Instance",
			ResourceIdentifiers: []vropsResourceIdentifier{
				{
					IdentifierType: vropsResourceIdentifierType{
						Name:               "VCURL",
						DataType:           "STRING",
						IsPartOfUniqueness: true,
					},
					Value: "vcenter.test.pdxfixit.com",
				},
			},
		},
		Description:                "test",
		CollectorID:                1,
		CollectorGroupID:           "foo",
		CredentialInstanceID:       "bar",
		MonitoringInterval:         2,
		NumberOfMetricsCollected:   3,
		NumberOfResourcesCollected: 4,
		LastHeartbeat:              5,
		LastCollected:              6,
		MessageFromAdapterInstance: "baz",
		Links:                      []vropsLink{},
		ID:                         "test-123",
	}

	records := []hostdb.Record{
		{
			ID:        "abc123",
			Type:      "test",
			Hostname:  "foo.pdxfixit.com",
			IP:        "10.20.30.40",
			Timestamp: "2019-01-01 00:00:00",
			Committer: "golang tests",
			Context:   map[string]interface{}{"test": true},
			Data:      []byte("{\"id\": \"abc123\", \"test\": \"wahoo\"}"),
			Hash:      "",
		}, {
			ID:        "def456",
			Type:      "test",
			Hostname:  "bar.pdxfixit.com",
			IP:        "10.50.60.70",
			Timestamp: "2019-01-01 00:00:00",
			Committer: "golang tests",
			Context:   map[string]interface{}{"test": true},
			Data:      []byte("{\"id\": \"def456\", \"test\": \"yahoo\"}"),
			Hash:      "",
		}, {
			ID:        "ghi789",
			Type:      "test",
			Hostname:  "baz.pdxfixit.com",
			IP:        "1.1.1.1",
			Timestamp: "2019-01-01 00:00:00",
			Committer: "golang tests",
			Context:   map[string]interface{}{"test": true},
			Data:      []byte("{\"id\": \"ghi789\", \"test\": \"zahoo\"}"),
			Hash:      "",
		},
	}

	recordSet := createRecordSet(adapter, records)

	assert.Equal(t, recordSet.Context["vc_name"], adapter.ResourceKey.Name, "vc_name")

	for _, identifier := range adapter.ResourceKey.ResourceIdentifiers {
		if identifier.IdentifierType.Name == "VCURL" {
			assert.Equal(t, recordSet.Context["vc_url"], identifier.Value, "vc_url")
			break
		}
	}

	if adapter.Description != "" {
		assert.Equal(t, recordSet.Context["vc_desc"], adapter.Description)
	}

	assert.Equal(t, recordSet.Type, strings.ToLower(fmt.Sprintf("vrops-%s", adapter.ResourceKey.AdapterKindKey)), "type")
	assert.NotEmpty(t, recordSet.Timestamp, "timestamp")
	assert.NotEmpty(t, recordSet.Context, "context")
	assert.Equal(t, recordSet.Committer, "hostdb-collector-vrops", "committer")
	assert.NotEmpty(t, recordSet.Records, "records")

}

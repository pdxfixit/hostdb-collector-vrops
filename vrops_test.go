package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdapterResources(t *testing.T) {

	// setup fake http server for test
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "{\"pageInfo\":{\"totalCount\":99,\"page\":1,\"pageSize\":10},\"links\":[{\"href\":\"foo\",\"rel\":\"SELF\",\"name\":\"current\"}],\"resourceList\":[{\"creationTime\":1524505047160,\"resourceKey\":{\"name\":\"vcenter.test.pdxfixit.com\",\"adapterKindKey\":\"TEST\",\"resourceKindKey\":\"TestAdapterInstance\",\"resourceIdentifiers\":[{\"identifierType\":{\"name\":\"TestName\",\"dataType\":\"STRING\",\"isPartOfUniqueness\":false},\"value\":\"vcenter.test.pdxfixit.com\"}]},\"resourceStatusStates\":[{\"adapterInstanceId\":\"3f6da67e-b49b-4714-963c-18b3b56e82a2\",\"resourceStatus\":\"DATA_RECEIVING\",\"resourceState\":\"STARTED\",\"statusMessage\":\"TESTING\"}],\"resourceHealth\":\"GREEN\",\"resourceHealthValue\":100.0,\"dtEnabled\":true,\"badges\":[{\"type\":\"TEST\",\"color\":\"GREEN\",\"score\":0.0}],\"relatedResources\":[],\"links\":[{\"href\":\"foo\",\"rel\":\"SELF\",\"name\":\"linkToSelf\"}],\"identifier\":\"2fb6adf9-7665-4bec-9d53-e49c5a71d63a\"}]}")
		if err != nil {
			t.Error(err.Error())
		}
	}))
	defer ts.Close()

	config.Vrops.Host = ts.URL

	adapterResources, err := getAdapterResources("fake-id-because-its-a-test", 0)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, adapterResources.PageInfo.TotalCount, 99, "PageInfo TotalCount")
	assert.Equal(t, adapterResources.PageInfo.Page, 1, "PageInfo Page")
	assert.Equal(t, adapterResources.PageInfo.PageSize, 10, "PageInfo PageSize")
	assert.Len(t, adapterResources.Links, 1, "Links count")
	assert.Equal(t, adapterResources.Links[0].Href, "foo", "Link Href")
	assert.Equal(t, adapterResources.Links[0].Rel, "SELF", "Link Rel")
	assert.Equal(t, adapterResources.Links[0].Name, "current", "Link Name")
	assert.Len(t, adapterResources.ResourceList, 1, "ResourceList count")
	assert.Equal(t, adapterResources.ResourceList[0].CreationTime, 1524505047160, "ResourceList CreationTime")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceKey.Name, "vcenter.test.pdxfixit.com", "ResourceList ResourceKey Name")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceKey.AdapterKindKey, "TEST", "ResourceList ResourceKey AdapterKindKey")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceKey.ResourceKindKey, "TestAdapterInstance", "ResourceList ResourceKey ResourceKindKey")
	assert.Len(t, adapterResources.ResourceList[0].ResourceKey.ResourceIdentifiers, 1, "ResourceList ResourceKey ResourceIdentifiers count")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceKey.ResourceIdentifiers[0].IdentifierType.Name, "TestName", "ResourceList ResourceKey ResourceIdentifier IdentifierType Name")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceKey.ResourceIdentifiers[0].IdentifierType.DataType, "STRING", "ResourceList ResourceKey ResourceIdentifier IdentifierType DataType")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceKey.ResourceIdentifiers[0].Value, "vcenter.test.pdxfixit.com", "ResourceList ResourceKey ResourceIdentifier Value")
	assert.Len(t, adapterResources.ResourceList[0].ResourceStatusStates, 1, "ResourceList ResourceStatusStates count")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceStatusStates[0].AdapterInstanceID, "3f6da67e-b49b-4714-963c-18b3b56e82a2", "ResourceList ResourceStatusStates AdapterInstanceID")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceStatusStates[0].ResourceStatus, "DATA_RECEIVING", "ResourceList ResourceStatusStates ResourceStatus")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceStatusStates[0].ResourceState, "STARTED", "ResourceList ResourceStatusStates ResourceState")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceStatusStates[0].StatusMessage, "TESTING", "ResourceList ResourceStatusStates StatusMessage")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceHealth, "GREEN", "ResourceList ResourceHealth")
	assert.Equal(t, adapterResources.ResourceList[0].ResourceHealthValue, float32(100), "ResourceList ResourceHealthValue")
	assert.Equal(t, adapterResources.ResourceList[0].DtEnabled, true, "ResourceList DtEnabled")
	assert.Len(t, adapterResources.ResourceList[0].Badges, 1, "ResourceList Badges count")
	assert.Equal(t, adapterResources.ResourceList[0].Badges[0].Type, "TEST", "ResourceList Badge Type")
	assert.Equal(t, adapterResources.ResourceList[0].Badges[0].Color, "GREEN", "ResourceList Badge Color")
	assert.Equal(t, adapterResources.ResourceList[0].Badges[0].Score, float32(0), "ResourceList Badge Score")
	assert.Len(t, adapterResources.ResourceList[0].RelatedResources, 0, "ResourceList RelatedResources")
	assert.Len(t, adapterResources.ResourceList[0].Links, 1, "ResourceList Links count")
	assert.Equal(t, adapterResources.ResourceList[0].Links[0].Href, "foo", "ResourceList Link Href")
	assert.Equal(t, adapterResources.ResourceList[0].Links[0].Rel, "SELF", "ResourceList Link Rel")
	assert.Equal(t, adapterResources.ResourceList[0].Links[0].Name, "linkToSelf", "ResourceList Link Name")
	assert.Equal(t, adapterResources.ResourceList[0].Identifier, "2fb6adf9-7665-4bec-9d53-e49c5a71d63a", "ResourceList Identifier")

}

func TestGetResourceProperties(t *testing.T) {

	data := `{"resourceId":"2fb6adf9-7665-4bec-9d53-e49c5a71d63a","property":[{"name":"test","value":"yes"}]}`

	// setup fake http server for test
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, data)
		if err != nil {
			t.Error(err.Error())
		}
	}))
	defer ts.Close()

	config.Vrops.Host = ts.URL
	config.Vrops.ResourceKindKeys = []string{"Test Adapter Instance"}

	testResources := []vropsResource{
		{
			1524505047160,
			vropsResourceKey{
				"vcenter.test.pdxfixit.com",
				"TEST",
				"Test Adapter Instance",
				[]vropsResourceIdentifier{
					{
						vropsResourceIdentifierType{
							"TestName",
							"STRING",
							false,
						},
						"vcenter.test.pdxfixit.com",
					},
				},
			},
			[]vropsResourceStatusState{
				{
					"3f6da67e-b49b-4714-963c-18b3b56e82a2",
					"DATA_RECEIVING",
					"STARTED",
					"TESTING",
				},
			},
			"GREEN",
			100.0,
			true,
			[]vropsBadge{
				{
					"TEST",
					"GREEN",
					0,
				},
			},
			[]interface{}{},
			[]vropsLink{},
			"2fb6adf9-7665-4bec-9d53-e49c5a71d63a",
		},
	}

	collection := getResourceProperties(testResources)

	assert.Len(t, collection, 1, "count of records")
	assert.Empty(t, collection[0].ID, "id should be empty")
	assert.Equal(t, collection[0].Type, "vrops-test-test_adapter_instance", "record type")
	assert.Empty(t, collection[0].Hostname, "hostname should be empty")
	assert.Empty(t, collection[0].IP, "ip should be empty")
	assert.NotEmpty(t, collection[0].Timestamp, "timestamp should not be empty")
	assert.NotEmpty(t, collection[0].Committer, "committer should not be empty")
	assert.Empty(t, collection[0].Context, "context should be empty")
	assert.Equal(t, collection[0].Data, json.RawMessage(data), "payload")
	assert.Empty(t, collection[0].Hash, "hash should be empty")

}

// todo: duplicate below test for -hostsystem records

// test the special function for type -virtualmachine records
func TestVirtualMachineMetadata(t *testing.T) {

	data := `{"resourceId":"01afa5ae-216f-4b27-91a7-43abfe5d5905","property":[{"name":"summary|guest|hostName","value":"localhost"},{"name":"summary|guest|ipAddress","value":"127.0.0.1"}]}`

	// setup fake http server for test
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, data)
		if err != nil {
			t.Error(err.Error())
		}
	}))
	defer ts.Close()

	config.Vrops.Host = ts.URL
	config.Vrops.ResourceKindKeys = []string{"virtualmachine"}

	testResources := []vropsResource{
		{
			1524505047160,
			vropsResourceKey{
				"vcenter.test.pdxfixit.com",
				"vmware",
				"virtualmachine",
				[]vropsResourceIdentifier{
					{
						vropsResourceIdentifierType{
							"TestName",
							"STRING",
							false,
						},
						"vcenter.test.pdxfixit.com",
					},
				},
			},
			[]vropsResourceStatusState{
				{
					"3f6da67e-b49b-4714-963c-18b3b56e82a2",
					"DATA_RECEIVING",
					"STARTED",
					"TESTING",
				},
			},
			"GREEN",
			100.0,
			true,
			[]vropsBadge{
				{
					"TEST",
					"GREEN",
					0,
				},
			},
			[]interface{}{},
			[]vropsLink{},
			"01afa5ae-216f-4b27-91a7-43abfe5d5905",
		},
	}

	collection := getResourceProperties(testResources)

	assert.Len(t, collection, 1, "count of records")
	assert.Empty(t, collection[0].ID, "id should be empty")
	assert.Equal(t, collection[0].Type, "vrops-vmware-virtualmachine", "record type")
	assert.Empty(t, collection[0].Hostname, "hostname should be empty")
	assert.Empty(t, collection[0].IP, "ip should be empty")
	assert.NotEmpty(t, collection[0].Timestamp, "timestamp should not be empty")
	assert.NotEmpty(t, collection[0].Committer, "committer should not be empty")
	assert.Empty(t, collection[0].Context, "context should be empty")
	assert.Equal(t, collection[0].Data, json.RawMessage(data), "payload")
	assert.Empty(t, collection[0].Hash, "hash should be empty")

}

func TestGetSessionToken(t *testing.T) {

	// setup fake http server for test
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "{\"token\":\"c0e7aa16-43e1-45a9-abc6-b90c8155a3a7::8066a69b-f6f9-4d65-84b4-691a13f1346a\",\"validity\":1546127294284,\"expiresAt\":\"Tuesday, January 1, 2019 0:00:00 AM UTC\",\"roles\":[]}")
		if err != nil {
			t.Error(err.Error())
		}
	}))
	defer ts.Close()

	config.Vrops.Host = ts.URL

	assert.NotEmpty(t, getSessionToken())

}

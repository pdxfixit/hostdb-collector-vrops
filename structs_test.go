package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVropsAdapterList_LoadFrom(t *testing.T) {

	// setup fake http server for test
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "{\"adapterInstancesInfoDto\":[{\"resourceKey\":{\"name\":\"Test Adapter\",\"adapterKindKey\":\"TEST\",\"resourceKindKey\":\"TestAdapterInstance\",\"resourceIdentifiers\":[{\"identifierType\":{\"name\":\"AUTODISCOVERY\",\"dataType\":\"STRING\",\"isPartOfUniqueness\":false},\"value\":\"false\"},{\"identifierType\":{\"name\":\"CLOUD_TYPE\",\"dataType\":\"STRING\",\"isPartOfUniqueness\":false},\"value\":\"TEST_CLOUD\"},{\"identifierType\":{\"name\":\"VCURL\",\"dataType\":\"STRING\",\"isPartOfUniqueness\":true},\"value\":\"vrops.test.pdxfixit.com\"}]},\"description\":\"Test Adapter Instance\",\"collectorId\":123,\"collectorGroupId\":\"27a89d5d-ab70-421d-aa42-d507e0540743\",\"credentialInstanceId\":\"f81e36ed-4f33-4007-a053-fe7478ab84f2\",\"monitoringInterval\":2,\"numberOfMetricsCollected\":1234,\"numberOfResourcesCollected\":42,\"lastHeartbeat\":1546909615138,\"lastCollected\":1546909516780,\"messageFromAdapterInstance\":\"Tests Started.\",\"links\":[{\"href\":\"/suite-api/api/adapters/15a4759d-0b2f-4432-bbfd-9a6f4cfab7e4\",\"rel\":\"SELF\",\"name\":\"linkToSelf\"},{\"href\":\"/suite-api/api/credentials/f81e36ed-4f33-4007-a053-fe7478ab84f2\",\"rel\":\"RELATED\",\"name\":\"linkToCredential\"}],\"id\":\"15a4759d-0b2f-4432-bbfd-9a6f4cfab7e4\"}]}")
		if err != nil {
			t.Error(err.Error())
		}
	}))
	defer ts.Close()

	adapterList := vropsAdapterList{}

	if err := adapterList.LoadFrom(ts.URL); err != nil {
		t.Errorf("%v", err)
	}

	// examine object
	assert.Len(t, adapterList.Instances, 1)
	assert.Equal(t, adapterList.Instances[0].ResourceKey.Name, "Test Adapter", "ResourceKey Name")
	assert.Equal(t, adapterList.Instances[0].ResourceKey.AdapterKindKey, "TEST", "ResourceKey AdapterKindKey")
	assert.Equal(t, adapterList.Instances[0].ResourceKey.ResourceKindKey, "TestAdapterInstance", "ResourceKey ResourceKindKey")
	assert.Len(t, adapterList.Instances[0].ResourceKey.ResourceIdentifiers, 3, "ResourceKey ResourceIdentifiers count")
	assert.Equal(t, adapterList.Instances[0].ResourceKey.ResourceIdentifiers[2].IdentifierType.Name, "VCURL", "ResourceKey ResourceIdentifier Name")
	assert.Equal(t, adapterList.Instances[0].ResourceKey.ResourceIdentifiers[2].IdentifierType.DataType, "STRING", "ResourceKey ResourceIdentifier DataType")
	assert.Equal(t, adapterList.Instances[0].ResourceKey.ResourceIdentifiers[2].Value, "vrops.test.pdxfixit.com", "ResourceKey ResourceIdentifier Value")
	assert.Equal(t, adapterList.Instances[0].Description, "Test Adapter Instance", "Description")
	assert.Equal(t, adapterList.Instances[0].CollectorID, 123, "CollectorID")
	assert.Equal(t, adapterList.Instances[0].CollectorGroupID, "27a89d5d-ab70-421d-aa42-d507e0540743", "CollectorGroupID")
	assert.Equal(t, adapterList.Instances[0].CredentialInstanceID, "f81e36ed-4f33-4007-a053-fe7478ab84f2", "CredentialInstanceID")
	assert.Equal(t, adapterList.Instances[0].MonitoringInterval, 2, "MonitoringInterval")
	assert.Equal(t, adapterList.Instances[0].NumberOfMetricsCollected, 1234, "NumberOfMetricsCollected")
	assert.Equal(t, adapterList.Instances[0].NumberOfResourcesCollected, 42, "NumberOfResourcesCollected")
	assert.Equal(t, adapterList.Instances[0].LastHeartbeat, 1546909615138, "LastHeartbeat")
	assert.Equal(t, adapterList.Instances[0].LastCollected, 1546909516780, "LastCollected")
	assert.Equal(t, adapterList.Instances[0].MessageFromAdapterInstance, "Tests Started.", "MessageFromAdapterInstance")
	assert.Len(t, adapterList.Instances[0].Links, 2, "Links count")
	assert.Equal(t, adapterList.Instances[0].Links[0].Href, "/suite-api/api/adapters/15a4759d-0b2f-4432-bbfd-9a6f4cfab7e4", "Link Href")
	assert.Equal(t, adapterList.Instances[0].Links[0].Rel, "SELF", "Link Rel")
	assert.Equal(t, adapterList.Instances[0].Links[0].Name, "linkToSelf", "Link Name")
	assert.Equal(t, adapterList.Instances[0].ID, "15a4759d-0b2f-4432-bbfd-9a6f4cfab7e4", "ID")

	// TODO: Test failure condition

}

func TestVropsAdapterResources_LoadFrom(t *testing.T) {

	// setup fake http server for test
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "{\"pageInfo\":{\"totalCount\":99,\"page\":1,\"pageSize\":10},\"links\":[{\"href\":\"foo\",\"rel\":\"SELF\",\"name\":\"current\"}],\"resourceList\":[{\"creationTime\":1524505047160,\"resourceKey\":{\"name\":\"vcenter.test.pdxfixit.com\",\"adapterKindKey\":\"TEST\",\"resourceKindKey\":\"TestAdapterInstance\",\"resourceIdentifiers\":[{\"identifierType\":{\"name\":\"TestName\",\"dataType\":\"STRING\",\"isPartOfUniqueness\":false},\"value\":\"vcenter.test.pdxfixit.com\"}]},\"resourceStatusStates\":[{\"adapterInstanceId\":\"3f6da67e-b49b-4714-963c-18b3b56e82a2\",\"resourceStatus\":\"DATA_RECEIVING\",\"resourceState\":\"STARTED\",\"statusMessage\":\"TESTING\"}],\"resourceHealth\":\"GREEN\",\"resourceHealthValue\":100.0,\"dtEnabled\":true,\"badges\":[{\"type\":\"TEST\",\"color\":\"GREEN\",\"score\":0.0}],\"relatedResources\":[],\"links\":[{\"href\":\"foo\",\"rel\":\"SELF\",\"name\":\"linkToSelf\"}],\"identifier\":\"2fb6adf9-7665-4bec-9d53-e49c5a71d63a\"}]}")
		if err != nil {
			t.Error(err.Error())
		}
	}))
	defer ts.Close()

	adapterResources := vropsAdapterResources{}

	if err := adapterResources.LoadFrom(ts.URL); err != nil {
		t.Errorf("%v", err)
	}

	// examine object
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

func TestVropsResourceProperties_LoadFrom(t *testing.T) {

	// setup fake http server for test
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "{\"resourceId\":\"2fb6adf9-7665-4bec-9d53-e49c5a71d63a\",\"property\":[{\"name\":\"test\",\"value\":\"yes\"}]}")
		if err != nil {
			t.Error(err.Error())
		}

	}))
	defer ts.Close()

	resourceProperties := vropsResourceProperties{}

	if err := resourceProperties.LoadFrom(ts.URL); err != nil {
		t.Errorf("%v", err)
	}

	// examine object
	assert.Equal(t, resourceProperties.ResourceID, "2fb6adf9-7665-4bec-9d53-e49c5a71d63a", "ResourceID")
	assert.Len(t, resourceProperties.Property, 1, "Property count")
	assert.Equal(t, resourceProperties.Property[0].Name, "test", "Property Name")
	assert.Equal(t, resourceProperties.Property[0].Value, "yes", "Property Value")

}

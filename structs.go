package main

import (
	"fmt"
	"log"
)

/*
	debug:       false
	sample_data: false
*/
type collectorConfig struct {
	Debug      bool `mapstructure:"debug"`
	SampleData bool `mapstructure:"sample_data"`
}

/*
	collector: {}
	vrops:     {}
*/
type globalConfig struct {
	Collector collectorConfig `mapstructure:"collector"`
	Vrops     vropsConfig     `mapstructure:"vrops"`
}

/*
	resourceKey:                {}
	description:                fancy
	collectorId:                7
	collectorGroupId:           27989d5d-ab70-421d-aa42-d50750540743
	credentialInstanceId:       f81d36ed-4f33-4007-a053-fe74786b84f2
	monitoringInterval:         5
	numberOfMetricsCollected:   7072
	numberOfResourcesCollected: 55
	lastHeartbeat:              1546909695138
	lastCollected:              1546909506780
	messageFromAdapterInstance: Trust Established.
	links:                      []
	id:                         1584759d-0b2f-4432-bbfd-9a6f4cfab764
*/
type vropsAdapterInstance struct {
	ResourceKey                vropsResourceKey `json:"resourceKey"`
	Description                string           `json:"description"`
	CollectorID                int              `json:"collectorId"`
	CollectorGroupID           string           `json:"collectorGroupId"`
	CredentialInstanceID       string           `json:"credentialInstanceId"`
	MonitoringInterval         int              `json:"monitoringInterval"`
	NumberOfMetricsCollected   int              `json:"numberOfMetricsCollected"`
	NumberOfResourcesCollected int              `json:"numberOfResourcesCollected"`
	LastHeartbeat              int              `json:"lastHeartbeat"`
	LastCollected              int              `json:"lastCollected"`
	MessageFromAdapterInstance string           `json:"messageFromAdapterInstance"`
	Links                      []vropsLink      `json:"links"`
	ID                         string           `json:"id"`
}

/*
	adapterInstancesInfoDto: []
*/
type vropsAdapterList struct {
	Instances []vropsAdapterInstance `json:"adapterInstancesInfoDto"`
}

func (obj *vropsAdapterList) LoadFrom(url string) (err error) {

	if config.Collector.Debug {
		log.Println(fmt.Sprintf(
			"Populating vropsAdapterList from %s.",
			url,
		))
	}

	if err := requestToStruct(url, &obj); err != nil {
		return err
	}

	if config.Collector.Debug {
		log.Println(fmt.Sprintf("%v", *obj))
	}

	return nil

}

/*
	pageInfo:     {}
	links:        []
	resourceList: []
*/
type vropsAdapterResources struct {
	PageInfo     vropsPageInfo   `json:"pageInfo"`
	Links        []vropsLink     `json:"links"`
	ResourceList []vropsResource `json:"resourceList"`
}

func (obj *vropsAdapterResources) LoadFrom(url string) (err error) {

	if config.Collector.Debug {
		log.Println(fmt.Sprintf(
			"Populating vropsAdapterResources from %s.",
			url,
		))
	}

	if err := requestToStruct(url, &obj); err != nil {
		return err
	}

	if config.Collector.Debug {
		log.Println(fmt.Sprintf("%v", *obj))
	}

	return nil

}

/*
	type:  RISK
	color: YELLOW
	score: 25.0
*/
type vropsBadge struct {
	Type  string  `json:"type"`
	Color string  `json:"color"`
	Score float32 `json:"score"`
}

/*
	host:             https://vrops.pdxfixit.com
	pageSize:         1000
	pass:             password
	resourceKindKeys: [ ClusterComputeResource Datastore VirtualMachine ]
*/
type vropsConfig struct {
	Host             string   `mapstructure:"host"`
	PageSize         int      `mapstructure:"pageSize"`
	Pass             string   `mapstructure:"pass"`
	ResourceKindKeys []string `mapstructure:"resourceKindKeys"`
	User             string   `mapstructure:"user"`
}

/*
	href: /suite-api/api/resources/2fb64df9-7665-4bec-9d53-e49c5a71563a
	rel:  SELF
	name: linkToSelf
*/
type vropsLink struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Name string `json:"name"`
}

/*
	totalCount: 63
	page:       0
	pageSize:   1000
*/
type vropsPageInfo struct {
	TotalCount int `json:"totalCount"`
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
}

/*
	name:  config|cpuAllocation|limit,
	value: -1.0
*/
type vropsProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
	creationTime:         1524505047760
	resourceKey:          {}
	resourceStatusStates: []
	resourceHealth:       GREEN
	resourceHealthValue:  100.0
	dtEnabled:            true
	badges:               []
	relatedResources:     []
	links:                []
	identifier:           2fb64df9-7665-4bec-9d53-e49c5a71563a
*/
type vropsResource struct {
	CreationTime         int                        `json:"creationTime"`
	ResourceKey          vropsResourceKey           `json:"resourceKey"`
	ResourceStatusStates []vropsResourceStatusState `json:"resourceStatusStates"`
	ResourceHealth       string                     `json:"resourceHealth"`
	ResourceHealthValue  float32                    `json:"resourceHealthValue"`
	DtEnabled            bool                       `json:"dtEnabled"`
	Badges               []vropsBadge               `json:"badges"`
	RelatedResources     []interface{}              `json:"relatedResources"`
	Links                []vropsLink                `json:"links"`
	Identifier           string                     `json:"identifier"`
}

/*
	identifierType: {}
	value:          vcenter.pdxfixit.com
*/
type vropsResourceIdentifier struct {
	IdentifierType vropsResourceIdentifierType `json:"identifierType"`
	Value          string                      `json:"value"`
}

/*
	name:               VMEntityName
	dataType:           STRING
	isPartOfUniqueness: false
*/
type vropsResourceIdentifierType struct {
	Name               string `json:"name"`
	DataType           string `json:"dataType"`
	IsPartOfUniqueness bool   `json:"isPartOfUniqueness"`
}

/*
	name:                vcenter.pdxfixit.com
	adapterKindKey:      VMWARE
	resourceKindKey:     VirtualMachine
	resourceIdentifiers: []
*/
type vropsResourceKey struct {
	Name                string                    `json:"name"`
	AdapterKindKey      string                    `json:"adapterKindKey"`
	ResourceKindKey     string                    `json:"resourceKindKey"`
	ResourceIdentifiers []vropsResourceIdentifier `json:"resourceIdentifiers"`
}

/*
	resourceId: 2fb64df9-7665-4bec-9d53-e49c5a71563a
	property:   []
*/
type vropsResourceProperties struct {
	ResourceID string          `json:"resourceId"`
	Property   []vropsProperty `json:"property"`
}

func (obj *vropsResourceProperties) LoadFrom(url string) (err error) {

	if config.Collector.Debug {
		log.Println(fmt.Sprintf(
			"Populating vropsResourceProperties from %s.",
			url,
		))
	}

	if err := requestToStruct(url, &obj); err != nil {
		return err
	}

	if config.Collector.Debug {
		log.Println(fmt.Sprintf("%v", *obj))
	}

	return nil

}

/*
	adapterInstanceId: 3f6da672-b49b-4714-963c-18b3b56e8222
	resourceStatus:    DATA_RECEIVING
	resourceState:     STARTED
	statusMessage:     Started.
*/
type vropsResourceStatusState struct {
	AdapterInstanceID string `json:"adapterInstanceId"`
	ResourceStatus    string `json:"resourceStatus"`
	ResourceState     string `json:"resourceState"`
	StatusMessage     string `json:"statusMessage"`
}

/*
	token:     c0e7aa16-43e1-4519-abc6-b90c8155a347::8066669b-f6f9-4d65-84b4-691913f1346a
	validity:  1546927294284
	expiresAt: Tuesday, January 8, 2019 6:01:34 AM UTC
	roles:     []
*/
type vropsSessionToken struct {
	Token     string   `json:"token"`
	Validity  int      `json:"validity"`
	ExpiresAt string   `json:"expiresAt"`
	Roles     []string `json:"roles"`
}

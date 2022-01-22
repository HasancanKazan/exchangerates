package main

import "time"

type Serie struct {
	SerieCode           string    `json:"serieCode", omitempty`
	DatagroupCode       string    `json:"datagroupCode", omitempty`
	SerieName           string    `json:"serieName", omitempty`
	SerieNameEng        string    `json:"serieNameEng", omitempty`
	FrequencyStr        string    `json:"frequencyStr", omitempty`
	DefaultAggMethodStr string    `json:"defaultAggMethodStr", omitempty`
	DefaultAggMethod    string    `json:"defaultAggMethod", omitempty`
	Tag                 string    `json:"tag", omitempty`
	TagEng              string    `json:"tagEng", omitempty`
	Datasource          string    `json:"datasource", omitempty`
	DatasourceEng       string    `json:"datasourceEng", omitempty`
	MetadataLink        string    `json:"metadataLink", omitempty`
	MetadataLinkEng     string    `json:"metadataLinkEng", omitempty`
	RevPolLink          string    `json:"revPolLink", omitempty`
	RevPolLinkEng       string    `json:"revPolLinkEng", omitempty`
	AppChaLink          string    `json:"appChaLink", omitempty`
	AppChaLinkEng       string    `json:"appChaLinkEng", omitempty`
	StartDate           time.Time `json:"startDate", omitempty`
	EndDate             time.Time `json:"endDate", omitempty`
}

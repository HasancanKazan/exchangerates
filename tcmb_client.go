package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	TCMBClient struct {
		baseUrl   string
		secretKey string
	}

	SerieTCMBResponse struct {
		SerieCode           string `json:"SERIE_CODE",omitempty`
		DatagroupCode       string `json:"DATAGROUP_CODE",omitempty`
		SerieName           string `json:"SERIE_NAME",omitempty`
		SerieNameEng        string `json:"SERIE_NAME_ENG",omitempty`
		FrequencyStr        string `json:"FREQUENCY_STR",omitempty`
		DefaultAggMethodStr string `json:"DEFAULT_AGG_METHOD_STR",omitempty`
		DefaultAggMethod    string `json:"DEFAULT_AGG_METHOD",omitempty`
		Tag                 string `json:"TAG",omitempty`
		TagEng              string `json:"TAG_ENG",omitempty`
		Datasource          string `json:"DATASOURCE",omitempty`
		DatasourceEng       string `json:"DATASOURCE_ENG",omitempty`
		MetadataLink        string `json:"METADATA_LINK",omitempty`
		MetadataLinkEng     string `json:"METADATA_LINK_ENG",omitempty`
		RevPolLink          string `json:"REV_POL_LINK",omitempty`
		RevPolLinkEng       string `json:"REV_POL_LINK_ENG",omitempty`
		AppChaLink          string `json:"APP_CHA_LINK",omitempty`
		AppChaLinkEng       string `json:"APP_CHA_LINK_ENG",omitempty`
		StartDate           string `json:"START_DATE",omitempty`
		EndDate             string `json:"END_DATE",omitempty`
	}
)

func (t *TCMBClient) GetSeries(ctx context.Context, code string) ([]SerieTCMBResponse, error) {
	url := fmt.Sprintf("%s/serieList/code=%s&key=%s", t.baseUrl, code, t.secretKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	var serieList []SerieTCMBResponse
	err = json.Unmarshal(body, &serieList)

	// for i, s := range serieList {
	// 	c := strings.Split(s.StartDate, "-")
	// 	x := c[2] + "-" + c[1] + "-" + c[0]
	// 	yourDate, _ := time.Parse("2006-01-02", x)

	// 	fmt.Println(yourDate, i, x)
	// }

	return serieList, err
}

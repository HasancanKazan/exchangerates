package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	TCMBRatesClient struct {
		StartDate string
		EndDate   string
		Type      string
		SecretKey string
		BaseUrl   string
	}

	TCMBRateResponseModel struct {
		TotalCount    int                      `json:"totalCount,omitempty"`
		CurrencyRates []map[string]interface{} `json:"items,omitempty"`
	}
)

func (t *TCMBRatesClient) GetRates(ctx context.Context, serieCodes string) (TCMBRateResponseModel, error) {
	var responseModel TCMBRateResponseModel
	url := fmt.Sprintf("%s/series=%s&startDate=%s&endDate=%s&type=%s&key=%s", t.BaseUrl, serieCodes, t.StartDate, t.EndDate, t.Type, t.SecretKey)
	response, err := http.Get(url)

	if err != nil {
		return responseModel, err
	}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	err = json.Unmarshal(body, &responseModel)
	return responseModel, err
}

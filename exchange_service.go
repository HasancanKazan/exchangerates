package main

import (
	"context"
	"time"
)

type ExchangeService struct {
	tcmbClient *TCMBClient
}

func (e ExchangeService) GetSeries(ctx context.Context, code string) ([]Serie, error) {
	res, err := e.tcmbClient.GetSeries(ctx, code)
	return ToSeries(res), err
}

func ToSeries(tcmbSeries []SerieTCMBResponse) (serie []Serie) {

	for idx, _ := range tcmbSeries {
		tcmbSeries := &tcmbSeries[idx]
		serie = append(serie, Serie{
			SerieCode:           tcmbSeries.SerieCode,
			DatagroupCode:       tcmbSeries.DatagroupCode,
			SerieName:           tcmbSeries.SerieName,
			SerieNameEng:        tcmbSeries.SerieNameEng,
			FrequencyStr:        tcmbSeries.FrequencyStr,
			DefaultAggMethodStr: tcmbSeries.DefaultAggMethodStr,
			DefaultAggMethod:    tcmbSeries.DefaultAggMethod,
			Tag:                 tcmbSeries.Tag,
			TagEng:              tcmbSeries.TagEng,
			Datasource:          tcmbSeries.Datasource,
			DatasourceEng:       tcmbSeries.DatasourceEng,
			MetadataLink:        tcmbSeries.MetadataLink,
			MetadataLinkEng:     tcmbSeries.MetadataLinkEng,
			RevPolLink:          tcmbSeries.RevPolLink,
			RevPolLinkEng:       tcmbSeries.RevPolLinkEng,
			AppChaLink:          tcmbSeries.AppChaLink,
			AppChaLinkEng:       tcmbSeries.AppChaLinkEng,
			StartDate:           time.Now(),
			EndDate:             time.Now(),
		})
	}

	return serie
}

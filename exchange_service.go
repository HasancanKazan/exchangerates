package main

import (
	"context"
	"strings"
	"time"
)

type ExchangeService struct {
	tcmbClient *TCMBClient
}

func (e ExchangeService) GetSeries(ctx context.Context, code string) (Serie, error) {
	res, err := e.tcmbClient.GetSeries(ctx, code)
	return ToSeries(res), err
}

func (e ExchangeService) GetAllSerieCodes(ctx context.Context, series Serie) string {
	var allSeriCodes string
	for _, serie := range series.Serie {
		allSeriCodes += serie.SerieCode + "-"
	}
	return strings.TrimRight(allSeriCodes, "-")
}

func ToSeries(tcmbSeries []SerieTCMBResponse) (serie Serie) {
	serie.DatagroupCode = tcmbSeries[0].DatagroupCode
	for idx, _ := range tcmbSeries {
		tcmbSeries := &tcmbSeries[idx]
		serie.Serie = append(serie.Serie, SerieDetail{
			SerieCode:           tcmbSeries.SerieCode,
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
			StartDate:           ConvertStringToTime(tcmbSeries.StartDate),
			EndDate:             ConvertStringToTime(tcmbSeries.EndDate),
		})
	}
	return serie
}

func ConvertStringToTime(date string) time.Time {
	sDate := strings.Split(date, "-")
	formattedDate := sDate[2] + "-" + sDate[1] + "-" + sDate[0]
	newDate, err := time.Parse(TIME_LAYOUT, formattedDate)
	if err != nil {
		return time.Now()
	}
	return newDate
}

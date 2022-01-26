package main

import (
	"context"
)

type RateService struct {
	rateClient *TCMBRatesClient
}

func (r RateService) GetCurrencyRates(ctx context.Context, serieCodes string) (TCMBRateResponseModel, error) {
	res, err := r.rateClient.GetRates(ctx, serieCodes)
	return res, err
}

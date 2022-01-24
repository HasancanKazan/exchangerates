package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	exchangeService ExchangeService
	rateService     RateService
}

func main() {
	app := echo.New()

	tcmbClient := &TCMBClient{
		baseUrl:   TCMB_BASE_URL,
		secretKey: TCMB_SECRET_KEY,
	}
	exchangeService := ExchangeService{tcmbClient}

	rateClient := &TCMBRatesClient{
		BaseUrl:   TCMB_BASE_URL,
		SecretKey: TCMB_SECRET_KEY,
		Type:      TYPE_JSON,
		StartDate: "24-01-2022",
		EndDate:   "24-01-2022",
	}
	rateService := RateService{rateClient}

	handler := Handler{
		exchangeService: exchangeService,
		rateService:     rateService,
	}
	app.GET("/series", handler.getExchangeRate)
	app.GET("/rates", handler.rates)
	app.Logger.Fatal(app.Start(":8000"))
}

func (h Handler) getExchangeRate(c echo.Context) error {
	res, err := h.exchangeService.GetSeries(c.Request().Context(), TCMB_DKDOV_CODE)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h Handler) rates(c echo.Context) error {
	res, err := h.exchangeService.GetSeries(c.Request().Context(), TCMB_DKDOV_CODE)
	if err != nil {
		return err
	}
	serieCodes := h.exchangeService.GetAllSerieCodes(c.Request().Context(), res)

	res2, err2 := h.rateService.GetSeries(c.Request().Context(), serieCodes)
	if err2 != nil {
		return err
	}
	return c.JSON(http.StatusOK, res2)
}

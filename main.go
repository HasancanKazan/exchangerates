package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	exchangeService ExchangeService
}

func main() {
	app := echo.New()
	tcmbClient := &TCMBClient{
		baseUrl:   TCMB_BASE_URL,
		secretKey: TCMB_SECRET_KEY,
	}
	exchangeService := ExchangeService{tcmbClient}
	handler := Handler{exchangeService: exchangeService}

	app.GET("/series", handler.getExchangeRate)
	app.Logger.Fatal(app.Start(":8000"))
}

func (h Handler) getExchangeRate(c echo.Context) error {
	res, err := h.exchangeService.GetSeries(c.Request().Context(), TCMB_DKDOV_CODE)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

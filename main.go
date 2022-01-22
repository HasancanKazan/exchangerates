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
	handler := Handler{
		exchangeService: exchangeService,
	}

	app.GET("/exchangerates", handler.getExchangeRate)
	app.Logger.Fatal(app.Start(":8000"))
}

func (h Handler) getExchangeRate(c echo.Context) error {
	res, err := h.exchangeService.GetSeries(c.Request().Context(), TCMB_DKDOV_CODE)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

// func convertToSeriesResponseModel(series *[]serieTCMB) int {
// 	serieList := new([]serie)

// 	serieList = series
// 	// for _, s := range *series {
// 	// 	c := strings.Split(s.StartDate, "-")
// 	// 	x := c[2] + "-" + c[1] + "-" + c[0]
// 	// 	yourDate, _ := time.Parse("2006-01-02", x)
// 	// 	serieList = append(serieList, )
// 	// }

// 	return 1
// }

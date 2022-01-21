package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/exchangerates", getExchangeRate)
	app.Logger.Fatal(app.Start(":8000"))
}

type serie struct {
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

func getExchangeRate(c echo.Context) error {

	// fmt.Print(TCMB_DATA_GROUPS_URL)

	response, err := http.Get("https://evds2.tcmb.gov.tr/service/evds/serieList/key=CyD61AWVNS&type=json&code=bie_dkdovytl")

	if err != nil {
		return err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	serieList := new([]serie)
	json.Unmarshal(body, &serieList)

	for i, s := range *serieList {
		c := strings.Split(s.StartDate, "-")
		x := c[2] + "-" + c[1] + "-" + c[0]
		yourDate, _ := time.Parse("2006-01-02", x)

		fmt.Println(yourDate, i, x)
	}

	i, _ := json.Marshal(serieList)

	return c.String(http.StatusAccepted, string(i))
}

func convertToSeriesResponseModel(series *[]serie) int {
	return 1
}

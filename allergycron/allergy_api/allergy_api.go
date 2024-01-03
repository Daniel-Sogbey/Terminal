package allergy_api

import (
	"allergycron/utils"
	"fmt"
	"net/http"
	"net/url"
)

type HourlyLoadResult struct {
	Total  int   `json:"total"`
	Hourly []int `json:"hourly"`
}

type HourlyLoadResponse struct {
	Success int              `json:"success"`
	Result  HourlyLoadResult `json:"result"`
}

type CurrentChartDataResult struct {
	Date    string  `json:"date"`
	Average float64 `json:"average"`
}

type CurrentChartDataResponse struct {
	Success int                      `json:"success"`
	Results []CurrentChartDataResult `json:"response"`
}

func GetHourlyLoadData() (*string, error) {

	queryParameters := url.Values{}
	queryParameters.Add("eID", "getHourlyLoadData")
	queryParameters.Add("type", "zip")
	queryParameters.Add("zip", "6800")
	queryParameters.Add("country", "AT")
	queryParameters.Add("pure_json", "1")

	response, err := utils.MakeHTTPRequest[HourlyLoadResponse]("https://www.polleninformation.at/index.php", http.MethodGet, nil, queryParameters, nil, HourlyLoadResponse{})

	if err != nil {
		return nil, err
	}

	averageLoad := 0
	for _, hour := range response.Result.Hourly {
		averageLoad += hour
	}

	averageLoad = averageLoad / len(response.Result.Hourly)

	scaleAverageLoad := averageLoad / 2

	formattedMessage := fmt.Sprintf("The average pollen load for today is %d", scaleAverageLoad)

	return &formattedMessage, nil
}

func GetCurrentChartData() (*string, error) {

}

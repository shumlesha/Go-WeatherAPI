package service

import (
	"WeatherfForecast/dto/forecast"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

const (
	BaseUrl     = "https://api.openweathermap.org/data/2.5"
	ApiKeyParam = "appid"
)

type WeatherClient interface {
	GetCurrentWeather(params map[string]string) (forecast.ForecastData, error)
}

type weatherClient struct {
	client *resty.Client
}

func NewWeatherClient(apiKey string) WeatherClient {
	client := resty.New()
	client.SetBaseURL(BaseUrl)
	client.SetQueryParam(
		ApiKeyParam, apiKey)

	return &weatherClient{
		client: client,
	}
}

func (w weatherClient) GetCurrentWeather(params map[string]string) (forecast.ForecastData, error) {

	var forecastData forecast.ForecastData
	_, err := w.client.R().
		SetQueryParams(params).
		SetResult(&forecastData).
		Get("/weather")

	logrus.Info("Weather data: ", forecastData)
	logrus.Info("Query: ", params)

	if err != nil {
		return forecastData, err
	}

	return forecastData, nil
}

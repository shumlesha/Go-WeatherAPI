package service

import (
	"WeatherfForecast/dto/forecast"
	"WeatherfForecast/pkg/repository"
)

type ForecastService interface {
	GetCurrentForecast(username string) (forecast.ForecastDto, error)
}

type forecastService struct {
	userRepository repository.UserRepository
	weatherClient  WeatherClient
}

func NewForecastService(userRepository repository.UserRepository, client WeatherClient) ForecastService {
	return &forecastService{
		userRepository: userRepository,
		weatherClient:  client,
	}
}

func (f forecastService) GetCurrentForecast(username string) (forecast.ForecastDto, error) {
	user, _ := f.userRepository.FindByUsername(username)

	cityName := user.City

	forecastData, err := f.weatherClient.GetCurrentWeather(
		map[string]string{
			"q": cityName,
		})

	return forecastData.ToForecastDto(), err
}

package forecast

import "fmt"

type ForecastData struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
}

func (f ForecastData) ToForecastDto() ForecastDto {
	return ForecastDto{
		Temperature: fmt.Sprintf("%.2f°C", f.Main.Temp-273.15),
		Humidity:    f.Main.Humidity,
		Summary:     fmt.Sprintf("%s: %s", f.Weather[0].Main, f.Weather[0].Description),
	}
}

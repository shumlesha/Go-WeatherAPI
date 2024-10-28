package forecast

type ForecastDto struct {
	Temperature string `json:"temperature"`
	Humidity    int    `json:"humidity"`
	Summary     string `json:"summary"`
}

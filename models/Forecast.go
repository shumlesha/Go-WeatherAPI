package models

import (
	"time"
)

type Forecast struct {
	BaseModel
	Date    time.Time `gorm:"not null" json:"date"`
	Temp    float64   `gorm:"not null" json:"temp"`
	Summary string    `gorm:"not null" json:"summary"`
}

package data

import "encoding/json"

type Pricing struct {
	ID           uint   `gorm:"primaryKey"`
	PlanName     string `gorm:"not null"`
	Description  string
	PricePerWeek float64   `gorm:"not null"`
	Features     []Feature `gorm:"many2many:pricing_features;"`
}

type Feature struct {
	ID          uint             `gorm:"primaryKey"`
	FeatureName string           `gorm:"not null;unique"`
	Metadata    *json.RawMessage `gorm:"type:json"`
}

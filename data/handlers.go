package data

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func GetPricing(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var pricing []Pricing
		db.Preload("Features").Find(&pricing)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(pricing)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

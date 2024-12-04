package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"proto_therapy_back/data"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&data.Pricing{}, &data.Feature{})

	if err != nil {
		panic("failed to migrate database")
	}

	data.SeedData(db)

	r := mux.NewRouter()

	r.HandleFunc("/pricing", data.GetPricing(db)).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

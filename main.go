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

	data.ResetDB(db)
	data.SeedData(db)

	r := mux.NewRouter()

	// Add CORS headers
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
			next.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/pricing", data.GetPricing(db)).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

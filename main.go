package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Country struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			url := "https://restcountries.com/v3.1/all"

			response, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			defer response.Body.Close()

			var countries []Country

			err = json.NewDecoder(response.Body).Decode(&countries)
			if err != nil {
				log.Fatal(err)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(countries)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 - Method Not Allowed"))
		}
	})

	log.Println("Listening on :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}


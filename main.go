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
	ID int `json:"id"`
}

var countries []Country

func main() {
	// Define a user map with usernames and passwords
	users := map[string]string{
		"admin": "password",
	}

	// Define a middleware that checks for a valid username and password
	authMiddleware := func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			username, password, ok := r.BasicAuth()

			if !ok || users[username] != password {
				w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password"`)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("401 - Unauthorized"))
				return
			}

			handler(w, r)
		}
	}

	// Create a new country
	http.HandleFunc("/countries", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var country Country
			err := json.NewDecoder(r.Body).Decode(&country)
			if err != nil {
				log.Fatal(err)
			}

			country.ID = len(countries) + 1
			countries = append(countries, country)

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Country created successfully"))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 - Method Not Allowed"))
		}
	}))

	// Read all countries
	http.HandleFunc("/countries", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(countries)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 - Method Not Allowed"))
		}
	}))

	// Update an existing country
	http.HandleFunc("/countries/{country_id}", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			countryID := r.URL.Path[len("/countries/"):]

			var country Country
			err := json.NewDecoder(r.Body).Decode(&country)
			if err != nil {
				log.Fatal(err)
			}

			for i, c := range countries {
				if c.ID == int(countryID) {
					countries[i] = country
					break
				}
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Country updated successfully"))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 - Method Not Allowed"))
		}
	}))

	// Delete an existing country
	http.HandleFunc("/countries/{country_id}", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			countryID := r.URL.Path[len("/countries/"):]

			for i, c := range countries {
				if c.ID == int(countryID) {
					countries = append(countries)
				}
			}
		}
	}))
}













	
				
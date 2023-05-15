

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

	http.HandleFunc("/", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
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
	}))

	log.Println("Listening on :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

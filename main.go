package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movies struct {
	Movies []Movie `json:"movies"`
}

type Movie struct {
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Imdb_id  string `json:"imdb_id"`
	Genre    string `json:"genre"`
	Language string `json:"language"`
	Size     int    `json:"size"`
	Quality  string `json:"quality"`
	Seeds    int    `json:"seeds"`
	Leeches  int    `json:"leeches"`
}

func main() {
	// Create a new gin.Engine object
	r := gin.Default()

	// Add your routes here
	r.GET("/movies", func(c *gin.Context) {
		// This line defines a new route that handles GET requests to the `/movies` path.

		// Make a request to the YTS API
		response, err := http.Get("https://yts.mx/api/v2/list_movies.json")
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Parse the JSON response
		var movies Movies
		err = json.Unmarshal(body, &movies)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Send the movies to the client
		c.JSON(200, movies)
	})

	// Start the server
	http.ListenAndServe(":8000", r)
}



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



// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Movie struct {
// 	Title       string   `json:"title"`
// 	Year        int      `json:"year"`
// 	Genres      []string `json:"genres"`
// 	Description string   `json:"description_full"`
// }

// type MoviesResponse struct {
// 	Data struct {
// 		Movies []Movie `json:"movies"`
// 	} `json:"data"`
// }

// func main() {
// 	r := gin.Default()
// 	r.Use(fetchMoviesMiddleware)

// 	r.GET("/movies", func(c *gin.Context) {
// 		movies, exists := c.Get("movies")
// 		if !exists {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, movies)
// 	})

// 	if err := r.Run(":8000"); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func fetchMoviesMiddleware(c *gin.Context) {
// 	response, err := http.Get("https://yts.mx/api/v2/list_movies.json")
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"message": fmt.Sprintf("Failed to fetch movies: %v", err),
// 		})
// 		return
// 	}
// 	defer response.Body.Close()

// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"message": fmt.Sprintf("Failed to fetch movies: %v", err),
// 		})
// 		return
// 	}

// 	var moviesResponse MoviesResponse
// 	err = json.Unmarshal(body, &moviesResponse)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"message": fmt.Sprintf("Failed to fetch movies: %v", err),
// 		})
// 		return
// 	}

// 	c.Set("movies", moviesResponse.Data.Movies)
// 	c.Next()
// }


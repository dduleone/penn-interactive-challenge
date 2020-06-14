package main

import (
	"fmt"
	"net/http"

	"dule.one/penn_interactive/imdb_challenge/api"

	"github.com/gin-gonic/gin"
)

const exposedPort = "80"

/*
Operation: 2itBush

Objectives:

[x] 1. Build 4 endpoints:
	[X] 1. GET /health
	[x] 2. GET /movies/tconst/:string
	[x] 3. GET /movies/startYear/:int
	[x] 4. GET /movies/genre/:string
[x] 2. When a request comes in:
	[x] 1. Open the TSV
	[x] 2. Read the contents
	[x] 3. Search for the answer (compile a list)
	[x] 4. Count results
	[X] 5. Format the results
	[X] 6. Send the response
[x] 3. Build Docker Container that will run this service
[x] 4. Build a Terraform Script to deploy it to AWS
[ ] 5. Capture output from 4 requests:
	[ ] 1. GET /health
	[ ] 2. GET /movies/tconst/tt0000492
	[ ] 3. GET /movies/startYear/2000
	[ ] 4. GET /movies/genre/drama
[ ] 6. Send in my results by 10pm Sunday

Notes:
	- I could speed this up by caching results to disk or redis/memcache
	- Also, I'd like to play with methods for buffering and returning results,
		to see if I can find a faster routine.
*/

func getGinEngine() *gin.Engine {
	router := gin.Default()

	/*
		Explicit 404 is unncessary. But APIs where parent routes clarify missing parameters are better for everyone.
		The assignment did not mention these routes, so I'm leaving them unimplemented and commented out.
	*/
	// router.GET("/", func(c *gin.Context) { c.String(http.StatusNotFound, "No such route.") })
	router.GET("/health", api.GetHealth)
	// router.GET("/movies", func(c *gin.Context) { c.String(http.StatusNotFound, "No such route.") })
	router.GET("/movies/tconst/:id", api.GetMoviesTconst)
	router.GET("/movies/startYear/:year", api.GetMoviesStartYear)
	router.GET("/movies/genre/:genre", api.GetMoviesGenre)

	fmt.Println("Server is online.")
	const x = http.StatusOK
	fmt.Println(x)
	return router
}

func main() {
	engine := getGinEngine()
	engine.Run(":" + exposedPort)
}

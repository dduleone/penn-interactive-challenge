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
[ ] 4. Build a Terraform Script to deploy it to AWS
[ ] 5. Capture output from 4 requests:
	[ ] 1. GET /health
	[ ] 2. GET /movies/tconst/tt0000492
	[ ] 3. GET /movies/startYear/2000
	[ ] 4. GET /movies/genre/drama
[ ] 6. Send in my results by 10pm Sunday

*/

func getGinEngine() *gin.Engine {
	router := gin.Default()

	// const epRoot = "/"
	const epHealth = "/health"
	// const epMoviesRoot = "/movies"
	const epMoviesTconst = "/movies/tconst/:id"
	const epMoviesStartYear = "/movies/startYear/:year"
	const epMoviesGenre = "/movies/genre/:genre"

	// router.GET(epRoot, api.GetRoot)
	router.GET(epHealth, api.GetHealth)
	// router.GET(epMoviesRoot, api.GetMoviesRoot)
	router.GET(epMoviesTconst, api.GetMoviesTconst)
	router.GET(epMoviesStartYear, api.GetMoviesStartYear)
	router.GET(epMoviesGenre, api.GetMoviesGenre)

	fmt.Println("Server is online.")
	const x = http.StatusOK
	fmt.Println(x)
	return router
}

func main() {
	engine := getGinEngine()
	engine.Run(":" + exposedPort)
}

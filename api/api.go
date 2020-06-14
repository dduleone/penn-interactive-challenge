package api

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type listResponse = struct {
	Count   int          `json:"count"`
	Results []ImdbRecord `json:"results"`
}

type healthcheckResponse = struct {
	Body string `json:"body"`
}

// Package and ship our response
func respond(c *gin.Context, results []ImdbRecord) {
	// If no results, return 404, because we're a well behaved API.
	httpcode := http.StatusOK
	if len(results) == 0 {
		httpcode = http.StatusNotFound
	}

	// Count our objects, and send them to the browser.
	response := listResponse{Count: len(results), Results: results}
	c.PureJSON(httpcode, response)
}

// I used this method to test against smaller, datasubsets.
func getDataFile() string {
	datafile := "./data/title.basics.tsv"
	// datafile := "./data/title.basics.head.tsv"
	// datafile := "./data/title.basics.head1000.tsv"
	// datafile := "./data/title.basics.head1000000.tsv"
	return datafile
}

// GetHealth GET /health
func GetHealth(c *gin.Context) {
	response := healthcheckResponse{Body: "OK"}
	c.PureJSON(http.StatusOK, response)
}

// GetMoviesTconst GET /movies/tconst/:id
func GetMoviesTconst(c *gin.Context) {
	// Grab URL Parameter and casefix it.
	idCasefixed := strings.ToLower(c.Param("id"))

	// Iterate dataset, and collect results.
	var results []ImdbRecord
	scanner, tsvFile := getFileScanner(getDataFile())
	for scanner.Scan() {
		oneLine := scanner.Text()
		fields := strings.Split(oneLine, "\t")
		oneRecord := arrayToRecord(fields)
		recordTconstCasefixed := strings.ToLower(oneRecord.Tconst)
		if recordTconstCasefixed == idCasefixed {
			results = append(results, oneRecord)
		}
	}

	// Clean up file handlers
	defer tsvFile.Close()

	// Ship the results
	respond(c, results)
}

// GetMoviesStartYear GET /movies/startYear/:year
func GetMoviesStartYear(c *gin.Context) {
	// Grab URL Parameter. We're expecting an int as a string, so no need to casefix.
	year := c.Params.ByName("year")

	// Iterate dataset, and collect results.
	var results []ImdbRecord
	scanner, tsvFile := getFileScanner(getDataFile())
	for scanner.Scan() {
		oneLine := scanner.Text()
		fields := strings.Split(oneLine, "\t")
		oneRecord := arrayToRecord(fields)
		if oneRecord.StartYear == year {
			results = append(results, oneRecord)
		}
	}

	// Clean up file handlers
	defer tsvFile.Close()

	// Ship the results
	respond(c, results)
}

// GetMoviesGenre GET /movies/genre/:genre
func GetMoviesGenre(c *gin.Context) {
	// Grab URL Parameter and make a casefixed version for comparisons
	genre := c.Params.ByName("genre")
	genreCasefixed := strings.ToLower(genre)

	// Iterate dataset, and collect results.
	var results []ImdbRecord
	scanner, tsvFile := getFileScanner(getDataFile())
	for scanner.Scan() {
		oneLine := scanner.Text()
		fields := strings.Split(oneLine, "\t")
		oneRecord := arrayToRecord(fields)

		// I don't like loops inside loops, but there shouldn't be more than a few genres per record.
		for _, oneGenre := range oneRecord.Genres {
			recordGenreCasefixed := strings.ToLower(oneGenre)
			if recordGenreCasefixed == genreCasefixed {
				results = append(results, oneRecord)
			}
		}
	}

	// Clean up file handlers
	defer tsvFile.Close()

	// Ship the results
	respond(c, results)
}

func getFileScanner(dataFile string) (*bufio.Scanner, *os.File) {
	// Open the datafile for reading.
	tsvFile, err := os.Open(dataFile)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}

	// Setup our scanner
	scanner := bufio.NewScanner(tsvFile)
	scanner.Split(bufio.ScanLines)

	// Return the tsvFile here, so I can close it on defer, in the parent.
	return scanner, tsvFile
}

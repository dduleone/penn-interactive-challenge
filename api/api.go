package api

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func return200(c *gin.Context, message string) {
	c.String(http.StatusOK, message)
}

func return404(c *gin.Context, message string) {
	c.String(http.StatusNotFound, message)
}

type listResponse = struct {
	Count   int          `json:"count"`
	Results []ImdbRecord `json:"results"`
}

type healthcheckResponse = struct {
	Body string `json:"body"`
}

// // GetRoot GET /
// func GetRoot(c *gin.Context) {
// 	return404(c, "GetRoot")
// }

// GetHealth GET /health
func GetHealth(c *gin.Context) {
	response := healthcheckResponse{Body: "OK"}
	c.PureJSON(http.StatusOK, response)
}

// // GetMoviesRoot GET /movies
// func GetMoviesRoot(c *gin.Context) {
// 	return404(c, "GetMoviesRoot")
// }

//http://localhost:8081/movies/tconst/tt10054590

func getDataFile() string {
	datafile := "./data/title.basics.tsv"
	// datafile := "./data/title.basics.tail.tsv"
	// datafile := "./data/title.basics.head.tsv"
	// datafile := "./data/title.basics.head1000.tsv"
	// datafile := "./data/title.basics.head1000000.tsv"
	return datafile
}

// GetMoviesTconst GET /movies/tconst/:id
func GetMoviesTconst(c *gin.Context) {
	id := c.Params.ByName("id")
	tsvFile, err := os.Open(getDataFile())
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer tsvFile.Close()

	scanner := bufio.NewScanner(tsvFile)
	scanner.Split(bufio.ScanLines)
	var allRecords []ImdbRecord

	for scanner.Scan() {
		oneLine := scanner.Text()

		fields := strings.Split(oneLine, "\t")
		oneRecord := arrayToRecord(fields)
		if oneRecord.Tconst == id {
			allRecords = append(allRecords, oneRecord)
		}
	}
	response := listResponse{Count: len(allRecords), Results: allRecords}
	c.PureJSON(http.StatusOK, response)
}

// GetMoviesStartYear GET /movies/startYear/:year
func GetMoviesStartYear(c *gin.Context) {
	year := c.Params.ByName("year")
	tsvFile, err := os.Open(getDataFile())
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer tsvFile.Close()

	scanner := bufio.NewScanner(tsvFile)
	scanner.Split(bufio.ScanLines)

	var allRecords []ImdbRecord
	for scanner.Scan() {
		oneLine := scanner.Text()

		fields := strings.Split(oneLine, "\t")
		oneRecord := arrayToRecord(fields)
		if oneRecord.StartYear == year {
			allRecords = append(allRecords, oneRecord)
		}
	}

	response := listResponse{Count: len(allRecords), Results: allRecords}
	c.PureJSON(http.StatusOK, response)

}

// GetMoviesGenre GET /movies/genre/:genre
func GetMoviesGenre(c *gin.Context) {
	genre := c.Params.ByName("genre")

	tsvFile, err := os.Open(getDataFile())
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer tsvFile.Close()

	scanner := bufio.NewScanner(tsvFile)
	scanner.Split(bufio.ScanLines)

	var allRecords []ImdbRecord

	for scanner.Scan() {
		oneLine := scanner.Text()

		fields := strings.Split(oneLine, "\t")
		oneRecord := arrayToRecord(fields)

		for _, oneGenre := range oneRecord.Genres {
			if oneGenre == genre {
				allRecords = append(allRecords, oneRecord)
			}
		}
	}

	response := listResponse{Count: len(allRecords), Results: allRecords}
	c.PureJSON(http.StatusOK, response)

	// jsonBytes, err := json.MarshalIndent(allRecords, "", "  ")

	// if err != nil {
	// 	return404(c, err.Error())
	// }

	// payload := string(jsonBytes)
	// return200(c, payload)
}

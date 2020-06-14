package api

import "strings"

func arrayToRecord(fields []string) ImdbRecord {
	var oneRecord ImdbRecord
	oneRecord.Tconst = fields[0]
	oneRecord.TitleType = fields[1]
	oneRecord.PrimaryTitle = fields[2]
	oneRecord.OriginalTitle = fields[3]
	oneRecord.IsAdult = strIntToBool(fields[4])
	oneRecord.StartYear = stringOrNull(fields[5])
	oneRecord.EndYear = stringOrNull(fields[6])
	oneRecord.RuntimeMinutes = stringOrNull(fields[7])
	oneRecord.Genres = strings.Split(fields[8], ",")
	return oneRecord
}

func stringOrNull(str string) string {
	if str == "\\N" {
		return ""
	}
	return str
}

func strIntToBool(strInt string) bool {
	if strInt == "1" {
		return true
	}
	return false
}

// ImdbRecord holds one record from the imdb:title.basics.tsv
type ImdbRecord struct {
	Tconst         string   `json:"tconst"`
	TitleType      string   `json:"titleType"`
	PrimaryTitle   string   `json:"primaryTitle"`
	OriginalTitle  string   `json:"originalTitle"`
	IsAdult        bool     `json:"isAdult"`
	StartYear      string   `json:"startYear,omitempty"`
	EndYear        string   `json:"endYear,omitempty"`
	RuntimeMinutes string   `json:"runtimeMinutes,omitempty"`
	Genres         []string `json:"genres"`
}

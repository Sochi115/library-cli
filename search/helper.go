package search

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func apiGetRequest(url string, responseObject interface{}) {
	msgString := fmt.Sprintf("Performing GET %s ...", url)
	writeToConsole(msgString)

	resp, err := http.Get(url)
	handleFatalErr(err)

	defer resp.Body.Close()

	writeToConsole("Decoding response JSON...")
	json.NewDecoder(resp.Body).Decode(responseObject)
}

func parseInputQuery(query string) string {
	tokens := strings.Split(strings.TrimSpace(query), " ")

	queryString := tokens[0]

	for _, token := range tokens[1:] {
		queryString += "+"
		queryString += token
	}

	return queryString
}

func handleFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func longStringToMultiline(text string) string {
	tokens := strings.Split(strings.TrimSpace(text), " ")
	result := ""
	for i := range tokens {
		result += tokens[i]
		result += " "
		if (i+1)%5 == 0 {
			result += "\n"
		}
	}
	return result
}

func writeToConsole(msg string) {
	io.WriteString(os.Stdout, msg)
	io.WriteString(os.Stdout, "\n")
}

func intSliceToString(arr []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ","), "[]")
}

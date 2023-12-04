package search

import (
	"fmt"
	"strings"
)

func parseInputQuery(query string) string {
	tokens := strings.Split(strings.TrimSpace(query), " ")

	queryString := tokens[0]

	for _, token := range tokens[1:] {
		queryString += "+"
		queryString += token
	}

	return queryString
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

func intSliceToString(arr []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ", "), "[]")
}

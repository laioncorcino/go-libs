package go_libs

import (
	"fmt"
	"strings"
)

// StringToMap recebe uma query string e transforma ela em um mapa
func StringToMap(queryString string) map[string]string {
	var pairs = strings.Split(queryString[1:], "&")
	queryStringMap := map[string]string{}

	for i := 0; i < len(pairs); i++ {
		var keysValue = strings.Split(pairs[i], "=")
		var keys = keysValue[0]
		var value = keysValue[1]
		queryStringMap[keys] = value
	}
	return queryStringMap
}

// MapToString recebe um mapa e transforma em uma query string
func MapToString(queryStringMap map[string]string) string {
	var queryString = "?"
	var index = 0
	for key, value := range queryStringMap {
		if index != 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%v=%v", key, value)
		index++
	}
	return queryString
}

package go_libs

import "strings"

func Analyze(queryString string) map[string]string {
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

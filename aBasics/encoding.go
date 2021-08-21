package aBasics

import "encoding/json"

type jsonExample struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func parseJson() string{
	jsonE := jsonExample{
		Name: "Go 'not golang' lang",
		Age:  25,
	}

	resultBytes, _ := json.Marshal(jsonE)
	return string(resultBytes)
}

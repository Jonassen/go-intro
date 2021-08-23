package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type jsonExample struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func parseJson() string{
	jsonE := jsonExample{
		Name: "Mats Jonassen",
		Age:  26,
	}

	utf8Bytes, _ := json.Marshal(jsonE)
	return string(utf8Bytes)
}

func bytesEverywhere() {
	var reader io.Reader = bytes.NewReader([]byte("Body text"))
	var _ io.Writer = bytes.NewBuffer([]byte("Buffer bytes"))

	req, _ := http.NewRequest(http.MethodPost, "bekk.no", reader)

	reader = req.Body
}

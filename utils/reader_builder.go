package utils

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func ReaderBuilder(url string) *strings.Reader{
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("error reading response body")
	}
	page := strings.NewReader(string(body))
	return page
}
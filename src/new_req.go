package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func req() {
	resp, err := http.Get("http://search.yahooapis.com/WebSearchService/V1/webSearch?appid=YahooDemo&query=finances&format=pdf")
	if err != nil {
		log.Panicln(err)
	}

	// Print HTTP Status
	fmt.Println(resp.Status)

	// Read display response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(body))
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	fmt.Print(json.NewDecoder(resp.Body))
	resp.Body.Close()
}

func main() {
	req()
}

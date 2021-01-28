package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	req, err := http.NewRequest("post", "http://httpbin.org/post", strings.NewReader(url.Values{}.Encode()))
	if err != nil {
		log.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:8888/")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(string(body))

}

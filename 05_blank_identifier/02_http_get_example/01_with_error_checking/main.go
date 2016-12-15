package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.goodlifefitness.com//assets/scripts/c-sharp/web-services/clubservice.svc/getclubsforcitypage")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()
	page, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", page)
}

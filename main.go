package main

import (
	"fmt"
	"log"
)

func main() {

	var request requester
	request.url = "http://ipinfo.io/"

	response, err := request.Get()
	if err != nil {
		log.Panic("FATAL: Cannot make request properly.")
	}
	fmt.Println(response)
}

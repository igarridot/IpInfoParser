package main

import (
	"fmt"
	"log"
)

func main() {

	var request requester
	response, err := request.Get("http://ipinfo.io/")

	if err != nil {
		log.Panic("FATAL: Cannot make request properly.")
	}
	fmt.Println(response)
}

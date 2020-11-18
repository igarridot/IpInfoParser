package main

import (
	"fmt"
	"log"
)

func main() {

	httpClientStruct := myhttp{}
	process(httpClientStruct)

}

func process(httpClientStruct MyHttp) {
	doMethod := mydo{}
	response, err := httpClientStruct.RequestBuilder(doMethod)
	if err != nil {
		log.Panic("The request has failed: ", err)
	}
	fmt.Println(response)
}

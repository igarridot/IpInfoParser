package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

//TODO: Try to manage Coordinates as [2]float64 from JSON Unmarshall
type IpProperties struct {
	IpAddr      string `json:"ip"`
	CityName    string `json:"city"`
	RegionName  string `json:"region"`
	CountryCode string `json:"country"`
	Coordinates string `json:"loc"`
	OrgName     string `json:"org"`
	PostalCode  int    `json:"postal,string"`
	TimeZone    string `json:"timezone"`
	ReadMe      string `json:"readme"`
}

func getIpInfo() (IpProperties, error) {

	var informationFomApi IpProperties

	requestToApi, err := http.Get("http://ipinfo.org")
	if err != nil {
		log.Println("ERROR: Cannot reach external resource http://ipinfo.org")
		return IpProperties{}, errors.New("Please check IInternet onnection.")
	}

	defer requestToApi.Body.Close()
	body, err := ioutil.ReadAll(requestToApi.Body)
	if err != nil {
		log.Println("ERROR: The request body is malformed.")
		return IpProperties{}, errors.New("Cannot parse request body from external resource properly.")
	}

	err = json.Unmarshal(body, &informationFomApi)
	if err != nil {
		log.Println("ERROR: Cannot parse response body to JSON")
		return IpProperties{}, errors.New("Some fields of the external resource has changed ot are not supported in this version.")
	}

	return informationFomApi, nil
}

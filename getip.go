package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type IpInformation struct {
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

type myHttpClient interface {
	Get() (IpInformation, error)
}

type requester struct{}

func (r *requester) Get(url string) (IpInformation, error) {

	var ipAttributes IpInformation
	httpClient := &http.Client{}
	buildRequest, err := http.NewRequest("GET", url, nil)
	resp, err := httpClient.Do(buildRequest)
	if err != nil {
		log.Println("We cannot reach the endpoint.", url)
		return IpInformation{}, errors.New("Incorrect GET")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("We cannot parse the JSON body", string(body))
		return IpInformation{}, errors.New("Incorrect JSON parse")
	}

	err = json.Unmarshal(body, &ipAttributes)
	if err != nil {
		log.Println("We cannot unmarshal body to struct.")
		return IpInformation{}, errors.New("Incorrect unmarshal")
	}
	return ipAttributes, nil
}

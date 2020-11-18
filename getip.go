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

type MyDo interface {
	Do(request *http.Request) (*http.Response, error)
}

type MyHttp interface {
	RequestBuilder(client MyDo) (IpInformation, error)
}

type mydo struct{}
type myhttp struct{}

func (d mydo) Do(request *http.Request) (*http.Response, error) {

	client := &http.Client{}
	return client.Do(request)
}

func (m myhttp) RequestBuilder(client MyDo) (IpInformation, error) {

	var ipAttributes IpInformation

	request, err := http.NewRequest("GET", "http://ipinfo.io/", nil)
	resp, err := client.Do(request)
	if err != nil {
		log.Println("We cannot reach the endpoint.")
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

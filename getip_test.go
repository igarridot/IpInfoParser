package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const jsonFixture = `{
  "ip": "127.0.0.1",
  "city": "Local",
  "region": "Host",
  "country": "LH",
  "loc": "0.0,-0.0",
  "org": "Localhost",
  "postal": "46000",
  "timezone": "Europe/Madrid",
  "readme": "https://ipinfo.io/missingauth"
}`

type HttpMockInterface struct {
	mock.Mock
}

func (t HttpMockInterface) Do(request *http.Request) (*http.Response, error) {
	args := t.Called(request)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestRequestBuilder(t *testing.T) {

	clientResponse := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonFixture))),
	}

	client := HttpMockInterface{}
	httpClientStruct := myhttp{}
	httpRequest, _ := http.NewRequest("GET", "http://ipinfo.io/", nil)
	client.On("Do", httpRequest).Return(clientResponse, nil)
	assert := require.New(t)
	returnedValue, _ := httpClientStruct.RequestBuilder(client)

	expectedOutput := IpInformation{
		IpAddr:      "127.0.0.1",
		CityName:    "Local",
		RegionName:  "Host",
		CountryCode: "LH",
		Coordinates: "0.0,-0.0",
		OrgName:     "Localhost",
		PostalCode:  46000,
		TimeZone:    "Europe/Madrid",
		ReadMe:      "https://ipinfo.io/missingauth",
	}

	assert.Equal(returnedValue, expectedOutput, "Test failed")
	client.AssertExpectations(t)
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Variables Globales

var token = ""
var server = "127.0.0.1"
var port = "5000"
var	client_id = ""
var	client_secret = ""

type ResponseToken struct {
	Type            string
	Username        string
	ApplicationName string
	ClientID        string
	TokenType       string
	AccessToken     string `json:"access_token"`
	Expires         int
	State           string
	Scope           string
}

type AtributosBusqueda struct {
	OriginLocationCode      string `form:"originLocationCode"`
	DestinationLocationCode string `form:"destinationLocationCode"`
	DepartureDate           string `form:"departureDate"`
	ReturnDate              string `form:"returnDate"`
	Adults                  int    `form:"adults"`
	IncludeAirlineCodes     string `form:"includedAirlineCodes"`
	NonStop                 bool   `form:"nonStop"`
	CurrencyCode            string `form:"currencyCode"`
	TravelClass             string `form:"travelClass"`
}

func obtenerToken(client_id string, client_secret string) string {

	client := &http.Client{}
	d := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", client_id, client_secret)
	var data = strings.NewReader(d)
	req, err := http.NewRequest("POST", "https://test.api.amadeus.com/v1/security/oauth2/token", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%s\n", bodyText)

	var response ResponseToken
	json.Unmarshal(bodyText, &response)

	return response.AccessToken
}

func buscarVuelos(c *gin.Context) {


	// Bind Query

	var url string

	var atributos AtributosBusqueda
	if c.ShouldBind(&atributos) == nil {
		url = fmt.Sprintf("https://test.api.amadeus.com/v2/shopping/flight-offers?originLocationCode=%s&destinationLocationCode=%s&departureDate=%s&adults=%s&includedAirlineCodes=EK&nonStop=true&currencyCode=CLP&travelClass=ECONOMY", atributos.OriginLocationCode, atributos.DestinationLocationCode, atributos.DepartureDate, atributos.Adults)
	}


	client := &http.Client{}
	// Preparamos la peticion
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	header_token := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", header_token)

	// Con client.Do(req) realizamos la peticion
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Aqui leemos el response body
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func main() {
	
	token = obtenerToken(client_id, client_secret)
	server_address := fmt.Sprintf("%s:%s", server, port)

	
	r := gin.Default(server_address)
	r.GET("/api/search", buscarVuelos)
	r.Run()
}

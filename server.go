package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

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

func buscarVuelos(token string) {

	client := &http.Client{}

	// Preparamos la peticion
	req, err := http.NewRequest("GET", "https://test.api.amadeus.com/v2/shopping/flight-offers?originLocationCode=DXB&destinationLocationCode=BKK&departureDate=2023-12-02&returnDate=2023-12-04&adults=1&includedAirlineCodes=EK&nonStop=true&currencyCode=CLP&travelClass=ECONOMY", nil)
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
	//buscarVuelos()
	token := obtenerToken("", "")
	fmt.Println(token)
	buscarVuelos(token)
}

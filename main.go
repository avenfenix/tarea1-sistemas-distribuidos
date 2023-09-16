package main

/////////////////////
/// goTravel MENU ///
/////////////////////

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/olekukonko/tablewriter"
)

type query struct {
	Origen   string
	Destino  string
	Fecha    string
	Cantidad string
}

type ResponseBusqueda struct {
	Data []DataObject `json:"data"`
}

type DataObject struct {
	Itineraries      []Itinerary       `json:"itineraries"`
	Price            Price             `json:"price"`
	TravelerPricings []TravelerPricing `json:"travelerPricings"`
}

type Itinerary struct {
	Duration string    `json:"duration"`
	Segments []Segment `json:"segments"`
}

type Segment struct {
	Departure       Airport `json:"departure"`
	Arrival         Airport `json:"arrival"`
	CarrierCode     string  `json:"carrierCode"`
	Number          string  `json:"number"`
	Duration        string  `json:"duration"`
	ID              string  `json:"id"`
	NumberOfStops   int     `json:"numberOfStops"`
	BlacklistedInEU bool    `json:"blacklistedInEU"`
}

type Airport struct {
	IATACode string `json:"iataCode"`
	Terminal string `json:"terminal"`
	At       string `json:"at"`
}

type Price struct {
	Currency   string `json:"currency"`
	Total      string `json:"total"`
	Base       string `json:"base"`
	GrandTotal string `json:"grandTotal"`
}

type TravelerPricing struct {
	TravelerID   string `json:"travelerId"`
	FareOption   string `json:"fareOption"`
	TravelerType string `json:"travelerType"`
	Price        Price  `json:"price"`
}

type Reserva struct{}

func menu_search(server string, port string) {
	var Query = query{}
	fmt.Printf("Aeropuerto de origen: ")
	// var origen string
	fmt.Scanln(&Query.Origen)
	fmt.Printf("Aeropuerto de destino: ")
	// var destino string
	fmt.Scanln(&Query.Destino)
	fmt.Printf("Fecha de salida: ")
	// var fecha string
	fmt.Scanln(&Query.Fecha)
	fmt.Printf("Cantidad de adultos: ")
	// var cantidad string
	fmt.Scanln(&Query.Cantidad)

	// ----------------------------------
	// BUSQUEDA CON API GOTRAVEL
	// ----------------------------------

	client := &http.Client{}

	// Preparamos el url para la peticion
	url := fmt.Sprintf("http://%s:%s/api/search?originLocationCode=%s&destinationLocationCode=%s&departureDate=%s&adults=%s&includedAirlineCodes=EK&nonStop=true&currencyCode=CLP&travelClass=ECONOMY", server, port, Query.Origen, Query.Destino, Query.Fecha, Query.Cantidad)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var response ResponseBusqueda
	json.NewDecoder(resp.Body).Decode(&response)

	fmt.Println(response.Data)

	// ----------------------------------
	// MOSTRAR TABLA CON RESULTADOS
	// ----------------------------------

	fmt.Printf("Se obtuvieron los siguientes resultados: \n")
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

/*
func getTokenAmadeus(){
	//
	url := "https://test.api.amadeus.com/v1/security/oauth2/token"

	body := []byte('grant_type=client_credentials&client_id={client_id}&client_secret={client_secret}')
} */

func main() {
	fmt.Println("Bievenido a goTravel!")
	in_menu := true
	for in_menu {
		fmt.Println("1. Realizar búsqueda.")
		fmt.Println("2. Obtener reserva.")
		fmt.Println("3. Salir.")
		fmt.Printf("Ingresa una opcion: ")
		var option string
		fmt.Scanln(&option)
		if option == "1" {
			menu_search("127.0.0.1", "5000")
		}
	}
}

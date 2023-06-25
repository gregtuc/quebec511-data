package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Worksite struct {
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
	Info string  `json:"info"`
}

var baseURL = "https://www.quebec511.info/en/Carte/Element.ashx?action="

func main() {
	// Get all municipal work sites in Montreal
	municipalSites, err := getCategory(baseURL + "ChantierMunicipal&xMin=-74.86661881860351&yMin=45.08690746124105&xMax=-72.45340318139648&yMax=45.95361168217988&lang=en&zoom=12")
	if err != nil {
		fmt.Println("Error in getting municipal work sites:", err)
	}


	// Get all minor provincial work sites in Montreal
	minorSites, err := getCategory(baseURL + "Chantier.Mineur&xMin=-74.86661881860351&yMin=45.08690746124105&xMax=-72.45340318139648&yMax=45.95361168217988&lang=en&zoom=12")
	if err != nil {
		fmt.Println("Error in getting minor provincial work sites:", err)
	}


	// Get all major provincial work sites in Montreal
	majorSites, err := getCategory(baseURL + "Chantier.Majeur&xMin=-74.86661881860351&yMin=45.08690746124105&xMax=-72.45340318139648&yMax=45.95361168217988&lang=en&zoom=12")
	if err != nil {
		fmt.Println("Error in getting major provincial work sites:", err)
	}
}

func getCategory(url string) ([]Worksite, error) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var worksites []Worksite
    err = json.NewDecoder(resp.Body).Decode(&worksites)
    if err != nil {
        fmt.Println("Error in decoding response body:", err)
		return nil, err
    }

	return worksites, nil
}
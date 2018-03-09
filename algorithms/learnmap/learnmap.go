package main

import (
	"fmt"
)

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	if captial, ok := countryCapitalMap["United States"]; ok {
		fmt.Println("Capital of United States is", captial)

	} else {
		fmt.Println("Capital of United States is not present")
	}

	countryCapitalMap1 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	delete(countryCapitalMap1, "France")

	for country := range countryCapitalMap1 {
		fmt.Println(countryCapitalMap1[country])
	}
}

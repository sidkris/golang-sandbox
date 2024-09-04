// this is similar to a class in Python
// each data point in a structure is called a field

package main

import "fmt"

func main() {

	type countries struct {
		country          string
		capitalCity      string
		approxPopulation int64
	}

	var countryData []countries

	countryData = append(countryData, countries{"India", "New Delhi", 1400000000})
	countryData = append(countryData, countries{"China", "Beijing", 1350000000})
	countryData = append(countryData, countries{"USA", "Washington D.C", 350000000})

	for _, country := range countryData {
		fmt.Printf("Country: %s, Capital: %s, Population: %d\n", country.country, country.capitalCity, country.approxPopulation)
	}
}

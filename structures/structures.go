// this is similar to a class in Python
// each data point in a structure is called a field

package main

import "fmt"

func main() {

	type Countries struct {
		country          string
		capitalCity      string
		approxPopulation int64
	}

	var countryData []Countries

	countryData = append(countryData, Countries{"India", "New Delhi", 1400000000})
	countryData = append(countryData, Countries{"China", "Beijing", 1350000000})
	countryData = append(countryData, Countries{"USA", "Washington D.C", 350000000})

	for _, country := range countryData {
		fmt.Printf("Country: %s, Capital: %s, Population: %d\n", country.country, country.capitalCity, country.approxPopulation)
	}
}

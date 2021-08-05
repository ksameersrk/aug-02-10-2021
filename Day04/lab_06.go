package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type City struct {
	name string
	country string
	population float64
}

type CountriesCitiesCollection map[string][]City
type CitiesCollection []City
var cities CitiesCollection = make(CitiesCollection, 0)
var countryCitiesMap CountriesCitiesCollection = make(CountriesCitiesCollection)

func main()  {
	loadCities()
	printCitiesGroupedByCountries()

	smallest, largest := printCityNameWithLargestAndSmallestPopulation()
	fmt.Println("City with smallest population", smallest)
	fmt.Println("City with largest population", largest)

	printPopulationOfChina()
}

func printPopulationOfChina()  {
	citiesInChina := countryCitiesMap["China"]
	var sum float64 = 0
	for i := 0; i < len(citiesInChina); i++ {
		sum += citiesInChina[i].population
	}
	fmt.Println(sum)
}

func printCityNameWithLargestAndSmallestPopulation() (smallest string, largest string) {
	max := cities[0].population
	min := cities[0].population
	for i := 0; i < len(cities); i++ {
		city := cities[i]
		if city.population > max {
			max = city.population
			largest = city.name
		}
		if city.population < min {
			min = city.population
			smallest = city.name
		}
	}
	return
}

func printCitiesGroupedByCountries()  {
	for i := 0; i < len(cities); i++ {
		city := cities[i]
		countryName := city.country
		if countryCitiesMap[countryName] == nil {
			cityArr := []City { city }
			countryCitiesMap[countryName] = cityArr
		} else {
			cityArr := countryCitiesMap[countryName]
			cityArr = append(cityArr, city)
			countryCitiesMap[countryName] = cityArr
		}
	}
	for country, cities := range countryCitiesMap {
		fmt.Println("************************* " + country + " *************************")
		for i := 0; i < len(cities); i++ {
			fmt.Println(cities[i].name)
		}
	}
}

func loadCities() {
	fileName := "./cities.csv"
	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.Contains(line, "City,Country,Population") {
				items := strings.Split(line, ",")
				cityName, countryName, population := items[0], items[1], items[2]
				populateCitiesCollection(cityName, countryName, population)
			}
		}
	}
}

func populateCitiesCollection(cityName string, countryName string, population string)  {
	populationInFloat, _ := strconv.ParseFloat(population, 64)
	city := City{name: cityName, country: countryName, population: populationInFloat}
	cities = append(cities, city)
}
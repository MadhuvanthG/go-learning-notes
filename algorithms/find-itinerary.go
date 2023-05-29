package main

import "fmt"

func findOrigin(itinerary map[string]string) (origin string) {
	sourceAirports := make([]string, 0, len(itinerary))
	reverseItinerary := make(map[string]string)
	for k, v := range itinerary {
		sourceAirports = append(sourceAirports, k)
		reverseItinerary[v] = k
	}

	for _, source := range sourceAirports {
		_, ok := reverseItinerary[source]
		if !ok {
			origin = source
			break
		}
	}

	return origin
}

func printItinerary(input map[string]string, src string) {
	dest, ok := input[src]
	if !ok {
		return
	}
	fmt.Printf("%s -> %s \n", src, dest)
	printItinerary(input, dest)
}

func runFindItinerary() {
	input := map[string]string{
		"HKG": "DXB",
		"FRA": "HKG",
		"DEL": "FRA",
	}

	input = map[string]string{
		"LAX": "DXB",
		"DFW": "JFK",
		"LHR": "DFW",
		"JFK": "LAX",
	}

	origin := findOrigin(input)
	printItinerary(input, origin)
}

package main

import "slices"

func getRawItineraryOnDay(line string, station int, forward bool, day int) []int {
	var output []int

	// Get itinerary
	it := getItinerary(line, station, forward)

	// Convert it
	workingIt := it[day]
	for hour, hourSpan := range workingIt {
		for _, minute := range hourSpan {
			output = append(output, hour*60+minute)
		}
	}

	return output
}

type lineStationCombo struct {
	line    string
	station int
	forward bool
}

type Departure struct {
	Line   string
	Hour   int
	Minute int
}

type departureTimeTable []Departure

func getDepartureTimeTable(day int, lines []lineStationCombo) departureTimeTable {
	var output departureTimeTable

	// Build unordered time-table
	for _, combo := range lines {
		this := getRawItineraryOnDay(combo.line, combo.station, combo.forward, day)
		for _, val := range this {
			output = append(output, Departure{
				Line:   combo.line,
				Hour:   val / 60,
				Minute: val % 60,
			})
		}
	}

	// Order time-table
	slices.SortFunc(output, func(a, b Departure) int {
		// First, compare hours
		if a.Hour != b.Hour {
			return a.Hour - b.Hour
		}
		// If hours are the same, compare minutes
		return a.Minute - b.Minute
	})

	return output
}

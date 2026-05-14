package main

import "fmt"

func main() {
	it := getDepartureTimeTable(0, []lineStationCombo{
		{
			line:    "8",
			station: 1,
			forward: true,
		},
		{
			line:    "7",
			station: 1,
			forward: true,
		},
		{
			line:    "410",
			station: 1,
			forward: true,
		},
	})

	fmt.Print("Time Table\n----------\n")
	for _, val := range it {
		fmt.Printf("Line %s [%02d:%02d]\n", val.Line, val.Hour, val.Minute)
	}
}

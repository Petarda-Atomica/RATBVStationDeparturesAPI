package main

import (
	"fmt"
	"time"
)

func styleTimeTable(input departureTimeTable) string {
	var table string
	table += "Bus,ETA,Plecare\n"

	now := time.Now().In(loc)
	nowMin := now.Hour()*60 + now.Minute()

	for _, val := range input {
		targetMin := val.Hour*60 + val.Minute
		ETA := targetMin - nowMin
		table += fmt.Sprintf("Linia %s,%d minute,%02d:%02d\n", val.Line, ETA, val.Hour, val.Minute)
	}

	return table
}

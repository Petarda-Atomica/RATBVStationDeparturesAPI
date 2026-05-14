package main

import (
	"log"
)

func main() {
	it := getItinerary("53", 1, false)

	log.Println(it[0][8])
}

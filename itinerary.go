package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type itinerary [7][24][]int

func getItinerary(line string, station int, forward bool) itinerary {
	var output itinerary

	c := colly.NewCollector()

	c.OnHTML("div#tabel2", func(h *colly.HTMLElement) {
		var week_rangeStart int
		var week_rangeEnd int
		var workingHour int

		h.DOM.Children().Each(func(i int, s *goquery.Selection) {
			id, _ := s.Attr("id")

			// Get the working range
			if id == "web_class_title" {
				startDescriptor, _ := removeDiacritics(strings.TrimSpace(strings.Split(s.Text(), "-")[0]))

				switch startDescriptor {
				case "LUNI":
					week_rangeStart = 0
				case "MARTI":
					week_rangeStart = 1
				case "MIERCURI":
					week_rangeStart = 2
				case "JOI":
					week_rangeStart = 3
				case "VINERI":
					week_rangeStart = 4
				case "SAMBATA":
					week_rangeStart = 5
				case "DUMINICA":
					week_rangeStart = 6
				}

				if strings.Contains(s.Text(), "-") {
					endDescriptor, _ := removeDiacritics(strings.TrimSpace(strings.Split(s.Text(), "-")[1]))
					switch endDescriptor {
					case "LUNI":
						week_rangeEnd = 0
					case "MARTI":
						week_rangeEnd = 1
					case "MIERCURI":
						week_rangeEnd = 2
					case "JOI":
						week_rangeEnd = 3
					case "VINERI":
						week_rangeEnd = 4
					case "SAMBATA":
						week_rangeEnd = 5
					case "DUMINICA":
						week_rangeEnd = 6
					}
				} else {
					week_rangeEnd = week_rangeStart
				}
			}

			// Get the working hour
			if id == "web_class_hours" {
				text := strings.TrimSpace(s.Text())
				if text == "Ora" {
					return
				}

				var err error
				workingHour, err = strconv.Atoi(text)
				if err != nil {
					log.Printf("Failed to get hour. Failed to decode string: %s\n", text)
				}
			}

			// Get the minute
			if id == "web_class_minutes" {
				text := strings.TrimSpace(s.Text())
				if text == "Minutul" || strings.Contains(text, "-") {
					return
				}

				s.Children().Each(func(i int, s *goquery.Selection) {
					id, _ := s.Attr("id")
					if id != "web_min" {
						return
					}

					min, err := strconv.Atoi(strings.TrimSpace(s.Text()))
					if err != nil {
						log.Printf("Failed to get minute. Failed to decode string: %s\n", s.Text())
						return
					}
					for i := week_rangeStart; i <= week_rangeEnd; i++ {
						output[i][workingHour] = append(output[i][workingHour], min)
					}
				})
			}
		})
	})

	// Get the url of the site we are trying to visit
	var direction string
	var linkIndex int
	if forward {
		direction = "dus"
		linkIndex = 2
	} else {
		direction = "intors"
		linkIndex = 1
	}
	visitUrl, err := url.JoinPath("https://www.ratbv.ro", "afisaje", fmt.Sprintf("%s-%s", line, direction), fmt.Sprintf("line_%s_%d_cl%d_ro.html", line, station, linkIndex))
	if err != nil {
		log.Fatal(err)
	}

	// Visit URL
	c.Visit(visitUrl)

	return output
}

package main

import (
	"errors"
	"fmt"
)

type Season int

const (
	SPRING Season = iota
	SUMMER
	AUTUMN
	WINTER
)

func monthToSeason(month int) (Season, error) {
	var season Season
	switch month {
	case 9, 10, 11:
		season = SPRING
	case 12, 1, 2:
		season = SUMMER
	case 3, 4, 5:
		season = AUTUMN
	case 6, 7, 8:
		season = WINTER
	default:
		return -1, errors.New("Invalid month")
	}
	return season, nil
}

func main() {
	month := 8
	season1, err := monthToSeason(month)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Month %d belongs to %d\n", month, season1)

}

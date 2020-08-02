package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	utcDT, _ := time.Parse(time.RFC3339, "2012-07-29T11:05:45+00:00")
	p("UTC dateTime: ", utcDT)
	tz, offset := utcDT.Zone()
	p("timezone: ", tz, offset)
	p("location", utcDT.Location())
	lc, _ := time.LoadLocation("Australia/Sydney")
	p(utcDT.In(lc).Format("2006-01-02T15:04:05.999999-07:00"))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type FruitBasket struct {
	Name    string
	Fruit   []string
	Id      int64  `json:"ref"`
	private string // An unexported field is not encoded.
	Created time.Time
}

func main() {
	jsonData := []byte(`{
		"Name":"Standard",
		"Fruit":["Apple","Banana","Orange"],
		"ref":999,
		"Created":"2020-07-30T08:34:41.19886+10:00",
		"other": 123
	}`)
	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", basket)
}

package main

import (
	"fmt"
	"time"
)

// defining two expensive computations
func compute1() string {
	// expensive computing
	time.Sleep(time.Second)
	return "3.14"
}
func compute2() string {
	// expensive computing
	time.Sleep(time.Second * 2)
	return "42"
}

func main() {

	go func() { r := compute1(); fmt.Println(r) }()

	time.Sleep(time.Second * 2)

}

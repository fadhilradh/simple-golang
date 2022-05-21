package main

import (
	"fmt"
	"log"
	"example.com/multifungsi"
)

func main() {
	log.SetPrefix("multifungsi : ")
	log.SetFlags(0)

	result, err := multifungsi.Greet("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result) 
}
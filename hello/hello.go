package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// fmt.Println(quote.Go())
	// fmt.Println(quote.Hello())
	// fmt.Println(quote.Glass())
	// fmt.Println(quote.Opt())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"World", "Go", "John Wick"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

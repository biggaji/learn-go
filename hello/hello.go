package main

import (
	"fmt"
	"log"

	"github.com/biggaji/greet"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	greeting, err := greet.SayHi("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)
}

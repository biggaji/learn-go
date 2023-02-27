package main

import (
	"fmt"
	"log"

	"github.com/biggaji/greet"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Tobi", "Anderson", "Gladys", "Dolapo", ""}

	greeting, err := greet.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting["Dolapo"])
}

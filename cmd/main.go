package main

import (
	"fmt"

	"github.com/alexvassiliou/gophercises/deck/deck"
)

func main() {
	d, _ := deck.New()

	for i, card := range d {
		fmt.Printf("card %d is the %s of %s", i+1, card.Type, card.Suit)
		fmt.Println()

	}

	deck.Shuffle(d)

	for i, card := range d {
		fmt.Printf("card %d is the %s of %s", i+1, card.Type, card.Suit)
		fmt.Println()

	}
}

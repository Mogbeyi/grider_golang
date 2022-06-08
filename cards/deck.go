package main

import (
	"fmt"
	"strings"
)

type (
	deck []string
	suit []string
	value []string
)

func newDeck() deck {
	cards := deck{}
	suits := suit{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := value{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value + " of " + suit)
		} 
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
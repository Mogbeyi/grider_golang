package main

import "fmt"

func main() {
	cards := newDeck() 
	// cards.print()	

	hand, remainingDeck := deal(cards, 7)
	hand.print()
	fmt.Println()
	remainingDeck.print()
}

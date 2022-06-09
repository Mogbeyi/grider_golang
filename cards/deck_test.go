package main

import(
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16 but go %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be ace of spades but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card to be four of clubs but got %v", d[len(d) - 1])
	}
} 

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck but got %v", len(loadedDeck))
	}
	
	os.Remove("_decktesting")
}
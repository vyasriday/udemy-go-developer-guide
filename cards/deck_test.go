package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16 but got %d", len(d))
	}

	if d[0] != "Jack of Clubs" {
		t.Errorf("Expected Jack of Clubs but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])

	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	cards := newDeck()
	saveToFile(cards, "_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

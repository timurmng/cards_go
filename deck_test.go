package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52 , got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, got %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	deck := newDeck()
	deck.saveToFile("_deckTesting", 0666)
	fileDeck := newDeckFromFile("_deckTesting")

	if len(fileDeck) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(fileDeck))
	}
	os.Remove("_deckTesting")
}

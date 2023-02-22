package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	newDeck := createDeck()

	if len(newDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(newDeck))
	}

	if newDeck[0] != "Ace of Spades ♠️" {
		t.Errorf("Expected first card was Ace of Spades ♠ but is %v", newDeck[0])
	}

	if newDeck[len(newDeck)-1] != "King of Clubs ♣️" {
		t.Errorf("Expected first card was King of Clubs ♣ but is %v", newDeck[len(newDeck)-1])
	}
}

func TestIODeckFile(t *testing.T) {

	newDeck := createDeck()

	newDeck.writeToFile("./_decktest.txt")

	deckLoaded := readFile("./_decktest.txt")

	if len(deckLoaded) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(newDeck))
	}

	err := os.Remove("./_decktest.txt")

	if err != nil {
		t.Fatal(err)
	}
}

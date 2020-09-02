package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of deck as 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected Ace of Spade, but got %v", d[0])
	}

	if (d[len(d)-1]) != "King of Club" {
		t.Errorf("Expected King of Club, but got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()

	d.saveToFile("_deckTesting")

	d = newDeckFromFile("_deckTesting")

	if len(d) != 52 {
		t.Errorf("Expected length of deck as 52, but got %d", len(d))
	}

	os.Remove("_deckTesting")
}

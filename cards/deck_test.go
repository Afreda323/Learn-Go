package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected length of deck: 52, Recieved length: %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card Two of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Clubs" {
		t.Errorf("Expected first card Two of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndDeckFromFile(t *testing.T) {
	fileName := "_deck_test_file.txt"
	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := deckFromFile(fileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of deck: 52, Recieved length: %v", len(loadedDeck))
	}

	os.Remove(fileName)
}

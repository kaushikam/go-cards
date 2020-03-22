package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected := 52
	if d.size() != expected {
		t.Errorf("Expected %d number of cards, but deck consists %d cards\n", expected, d.size())
	}

	firstCard := card{"Spades", "One"}
	if d[0] != firstCard {
		t.Errorf("Expected first card as %s, but got %s", firstCard, d[0])
	}

	lastCard := card{"Diamonds", "King"}
	if d[d.size()-1] != lastCard {
		t.Errorf("Expected last card as %s, but got %s", lastCard, d[d.size()-1])
	}
}

func TestSaveDeckToFile(t *testing.T) {
	filename := "test_deck.txt"
	os.Remove(filename)

	d := newDeck()
	err := d.saveToFile(filename)
	if err != nil {
		t.Errorf("File cannot be saved due to: %s", err.Error())
		os.Exit(1)
	}

	loaded := newDeckFromFile(filename)
	if loaded.size() != d.size() {
		t.Errorf("File loaded from disk contains %d number of cards, while we saved %d cards", loaded.size(), d.size())
	}

	os.Remove(filename)
}

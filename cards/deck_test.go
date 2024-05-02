package main

import (
	"os"
	"testing"
)

// t.Errorf("Expected deck length of 16, but got %v", len(d))

var numOfValues = 4
var numOfSuits = 4
var numOfCards = numOfValues * numOfSuits

func TestNewDeck(t *testing.T) {
	d := newDeck()
	// Check length
	dLen := len(d)
	if  dLen != numOfCards {
		t.Errorf("Expected deck length of %v, but got %v", numOfCards, dLen)
	}
	 
	// Check first card
	firstCard := d[0]
	expectedFirstCard := "Ace of Spades"
	if firstCard != expectedFirstCard {
		t.Errorf("Expected card %s, but got %s", firstCard, expectedFirstCard)
	}

	// Check last card
	lastCard := d[dLen - 1]
	expectedLastCard := "Four of Clubs"
	if lastCard != expectedLastCard {
		t.Errorf("Expected card %s, but got %s", lastCard, expectedLastCard)
	}
}

func TestDeckFiles(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	d2 := newDeckFromFile(filename)
	d2Len := len(d2)
	if  d2Len != numOfCards {
		t.Errorf("Expected deck length of %v, but got %v", numOfCards, d2Len)
	}

	os.Remove(filename)
}
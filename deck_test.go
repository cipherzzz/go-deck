package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	actualLength := len(deck)
	expectedLength := 56
	if actualLength != expectedLength {
		t.Errorf("Expected deck to be of length %d. Actual length was %d", expectedLength, actualLength)
	}

	expectedFirst := "Ace of Hearts"
	actualFirst := deck[0]
	if actualFirst != expectedFirst {
		t.Errorf("Expected first card to be %v. Actual card was %v", expectedFirst, actualFirst)
	}

	expectedLast := "King of Diamonds"
	actualLast := deck[len(deck)-1]
	if actualLast != expectedLast {
		t.Errorf("Expected last card to be %v. Actual card was %v", expectedLast, actualLast)
	}
}

func TestSaveDeckAndGetDeck(t *testing.T) {
	fileName := "_testdeck.txt"
	os.Remove(fileName)

	savedDeck := newDeck()
	err := saveDeck(fileName, savedDeck)
	if err != nil {
		t.Errorf("Unable to save deck to file: %v", fileName)
	}

	openedDeck, err := getDeck(fileName)
	if err != nil {
		os.Remove(fileName)
		t.Errorf("Unable to retrieve deck from file: %v", fileName)
	}

	if len(openedDeck) != 56 || len(savedDeck) != len(openedDeck) {
		t.Errorf("Expected opened deck to equal saved deck with a length of: %d", 56)
	}

	os.Remove(fileName)
}

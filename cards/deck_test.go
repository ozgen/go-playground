package main

import (
	"os"
	testing "testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected '%v' but got '%v'", 40, len(d))
	}

	if d[0] != "ace of hearts" {
		t.Errorf("Expected '%v'but got '%v'", "ace of hearts", d[0])
	}

	if d[len(d)-1] != "ten of clubs" {
		t.Errorf("Expected '%v'but got '%v'", "ten of clubs", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFileName := "_decktesting"
	os.Remove(testFileName)

	d := newDeck()

	d.saveToFile(testFileName)

	loadedDeck := newDeckFromFile(testFileName)

	if len(d) != len(loadedDeck) {
		t.Errorf("Expected '%v'but got '%v'", len(d), len(loadedDeck))
	}
	os.Remove(testFileName)
}

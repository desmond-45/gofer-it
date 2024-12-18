package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedValue := 16
	firstCard := "Ace of Spades"
	if len(d) != expectedValue {
		t.Errorf("Expected deck length of %d, but got %d", expectedValue, len(d))
	}
	if d[0] != firstCard {
		t.Errorf("Expected first card %s, but got %s", firstCard, d[0])
	}
}

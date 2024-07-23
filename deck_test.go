package main

import (
	"os"
	"testing"
)

func TestCreateNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 52{
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}

func TestSaveDeckToFile(t *testing.T){
	f :="_decktesting"
	os.Remove(f)

	d := newDeck()
	d.saveToFile(f)

	loadedDeck := newDeckFromFile(f)

	if len(loadedDeck) != len(d){
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}
	os.Remove(f)
}
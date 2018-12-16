package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

// Returns a list of playing cards
func newDeck() deck {
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jack", "Queen", "King", "Ace",
	}

	cards := deck{}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

// Iterates over cards and prints them to the console
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Returns a hand of cards,
// as well as updated deck
func deal(d deck, handSize int) (hand deck, deck deck) {
	return d[:handSize], d[handSize:]
}

// Return a string version of the deck
func (d deck) toString() string {
	return strings.Join(d, ", ")
}

// Saves the deck to a txt file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Create and return a deck of files from a file
// Exits process if error
func deckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ", ")
	return deck(ss)
}

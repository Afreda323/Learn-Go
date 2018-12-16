package main

import "fmt"

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

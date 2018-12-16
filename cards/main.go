package main

import "fmt"

func main() {
	deck := newDeck()
	deck.print()
	fmt.Println("____")
	deck.shuffle()
	deck.print()
}

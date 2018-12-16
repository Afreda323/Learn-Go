package main

func main() {
	cards := newDeck()
	hand, deck := deal(cards, 6)
	hand.print()
	deck.print()
}

package main

func main() {
	cards := newDeck()
	cards.shuffle()

	hand, deck := deal(cards, 6)

	hand.print()
	deck.print()
}

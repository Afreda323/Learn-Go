package main

func main() {
	cards := newDeck()
	hand, deck := deal(cards, 6)
	hand.saveToFile("hand.csv")
	deck.saveToFile("deck.csv")
}

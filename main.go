package main

func main() {
	cards := newDeck()
	hand, remianingDeck := deal(cards, 5)
	hand.print()
	remianingDeck.print()
}

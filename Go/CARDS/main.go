package main

func main() {
	card := newDeck()
	// card.print()
	// hand, remainingCards := deal(card, 5)

	// hand.print()
	// remainingCards.print(
	card.shuffle()
	card.print()
}

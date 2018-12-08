package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Fix of Spades")

	cards := deck{}
	cards = newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

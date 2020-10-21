package main

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, cards := cards.deal()
	// hand.print("hand")
	// cards.print("deck")
	saveDeck("hand.txt", hand)
	open, _ := getDeck("hand.txt")
	open.print("opened")
}

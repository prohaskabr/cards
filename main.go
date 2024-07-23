package main

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 2)

	// fmt.Println(hand.toString())
	// fmt.Println(remainingCards.toString())

	//cards.saveToFile("deck.txt")

	//newCards := newDeckFromFile("deck.txt")

	//newCards.print()

	cards.shuffle()
	cards.print()
}

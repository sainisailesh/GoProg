package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 26)
	// hand.print()
	// fmt.Println("----------------------------")
	// remainingCards.print()
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
}

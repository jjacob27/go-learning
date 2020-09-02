package main

func main() {
	//cards := newDeck()
	//hand, _ := deal(cards, 5)
	//hand.print()
	//fmt.Println(hand.toString())
	//hand.saveToFile("my_cards")

	// cards2 := newDeckFromFile("my_cards")
	// fmt.Println(cards2.toString())

	cards := newDeck()
	cards.shuffle()
	hand, _ := deal(cards,10)
	hand.print()

}

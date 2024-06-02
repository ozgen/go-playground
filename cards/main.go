package main

func main() {
	//var card string = "Ace of Spades"
	// go run main.go deck.go -> to run main with deck
	//card := "Test cards"
	//card = "Five of Diamonds"
	//newCard := createNewCardInt()
	//fmt.Println(card)
	//fmt.Println(newCard)
	//cartList := deck{"ace of diamonds", createNewCard()}
	//cartList = append(cartList, "six of spades")
	//fmt.Println(cartList)
	//for i, c := range cartList {
	//	fmt.Println(i, c)
	//}

	//cartList.print()

	cards := newDeck()
	//hand, remainingCards := deal(cards, 20)
	//hand.print()
	//remainingCards.print()
	//testString := "test string"
	//fmt.Println([]byte(testString))
	//fmt.Println(cards.toString())
	cards.saveToFile("card_file")
	existingCards := newDeckFromFile("card_file")
	existingCards.shuffle()
	existingCards.print()
}

func createNewCard() string {
	return "five of diamonds"
}
func createNewCardInt() int {
	return 12
}

package main

func main() {

	//function call
	//card := newCard()

	//slice manual appending
	//cards := deck{"Ace of Diamons", newCard()}
	//cards = append(cards, "Six of Spades")

	//fmt.Println(card)
	//fmt.Println(cards)

	cards := newDeck()
	cards.print()
}

//function to create new card
// func newCard() string {
// 	return "Five of Diamonds"
// }

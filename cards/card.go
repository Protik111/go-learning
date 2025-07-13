package main

import "fmt"

func main() {

	card := newCard()

	//slice
	cards := []string{"Ace of Diamons", newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println(card)
	//fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

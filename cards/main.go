package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	hand, remainingCards := deal(cards, 3)
	fmt.Println("HAND ----")
	hand.print()
	fmt.Println("REMAINING CARDS ----")
	remainingCards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("hello.txt")
	fmt.Println("CREATED DECK FROM FILE ----")
	newCreatedDeck := createDeckFromFile("hello.txt")
	newCreatedDeck.print()
	fmt.Println("SHUFFLE ------")
	newCreatedDeck.shuffle()
	newCreatedDeck.print()
}

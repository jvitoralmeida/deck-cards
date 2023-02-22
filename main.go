package main

import "fmt"

func main() {
	myDeck := createDeck()
	fmt.Println("----------------Deck before shuffle------------------")
	myDeck.print()
	fmt.Println("----------------Deck after shuffle-------------------")
	myDeck.shuffle()
	myDeck.print()
}

package main

import "fmt"

func main() {
	var deck = deck{}
	//fmt.Println(cards.toString())
	deck.shuffle()
	fmt.Println(deck.size())
}

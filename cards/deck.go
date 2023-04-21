package main

import "fmt"

// Create a new type of 'deck', which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// 用巢狀迴圈將各種花色與數值組成的牌張加入 cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

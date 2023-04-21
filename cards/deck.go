package main

import "fmt"

// Create a new type of 'deck', which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// 用巢狀迴圈將各種花色與數值組成的牌張加入 cards。由於本邏輯不需使用索引值，因此用 _ 取代 i, j 兩個變數。
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

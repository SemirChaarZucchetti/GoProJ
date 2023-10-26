package main

import "fmt"

/* crate a new type of "deck"
that is a slice of strings
*/

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {

			cards = append(cards, value+" of "+suit)

		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

package main

import "fmt"

/* crate a new type of "deck"
that is a slice of strings
*/

type Deck []string

func (d Deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

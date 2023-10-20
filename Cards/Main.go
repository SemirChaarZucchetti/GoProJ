package main

import "fmt"

func main() {

	cards := []string{"Ace of Diamond", newCard()}

	cards = append(cards, "Six of spade")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of diamond"
}

package main

func main() {

	cards := Deck{"Ace of Diamond", newCard()}

	cards = append(cards, "Six of spade")

	cards.print()
}

func newCard() string {
	return "Five of diamond"
}

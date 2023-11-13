package main

func main() {

	deck := newDeckFromFile("il mio mazzo")
	deck.shuffle()
	deck.saveFile("mazzo mischiato")

	hand, remainDeck := deal(newDeckFromFile("mazzo Mischiato"), 5)

	print("----------------------------------\n")
	hand.saveFile("la mia mano")
	hand.print()
	print("----------------------------------\n")

	print("----------------------------------\n")
	remainDeck.saveFile("mazzo restante")
	remainDeck.print()
	print("----------------------------------\n")

}

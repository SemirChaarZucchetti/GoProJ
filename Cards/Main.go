package main

import (
	"fmt"
	"os"
)

func main() {

	filesClean := []string{"la mia mano", "mazzo mischiato", "mazzo restante", "il mio mazzo"}

	for _, filePath := range filesClean {
		err := os.Remove(filePath)

		if err != nil {
			fmt.Println("Error deleting file:", err)
		} else {
			fmt.Println("File deleted:", filePath)
		}
	}

	deck := newDeck()
	deck.shuffle()
	deck.print()
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

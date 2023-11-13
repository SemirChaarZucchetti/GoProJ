package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

/* crate a new type of "deck"
that is a slice of strings
*/

type Deck []string

func (d Deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveFile(filename string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the deck as a string to the file
	_, err = io.WriteString(file, d.toString())
	if err != nil {
		return err
	}

	return nil
}

func newDeckFromFile(filename string) deck {
	// Read the file's content
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Split the content into slices and return it as a deck
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

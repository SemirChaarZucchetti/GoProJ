package main

import (
	"fmt"
)

type number []int

func (n number) evenOrOdd() {

	for _, numero := range n {
		if numero%2 == 0 {
			fmt.Println("%v is odd", numero)
		} else {
			fmt.Println("%v is even", numero)
		}
	}

}

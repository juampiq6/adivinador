package main

import (
	"fmt"
)

func humanoAdivina() {

	fmt.Print("\n_______________________________________________________________________________________________\n")
	fmt.Print("\n                                    Adivine el número:                                     ")
	fmt.Print("\n_______________________________________________________________________________________________\n")
	for {
		num, _ := preguntarNum()
		bien, reg := 
		if bien == 4 {
			fmt.Print("\n *** Felicitaciones, adivino el número! *** ")
			break
		} else {
			fmt.Print("BIEN : ", bien, " - REGULAR : ", reg)
		}
	}
}
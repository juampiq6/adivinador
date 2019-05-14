package main

import "fmt"

func main() {

	fmt.Print(" ¿ Quiere ser el pensador(p) o el adivinador(A) del número ? : ")
	var juego string
	fmt.Scan(&juego)
	if juego == "p" {
		humanoAdivina()
	} else {
		maquinaAdivina()
	}
}

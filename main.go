package main

import "fmt"

func main() {

	fmt.Print(" ¿ Quiere ser el pensador(P) o el adivinador(a) del número ? : ")
	var juego string
	fmt.Scan(&juego)
	if juego == "a" {
		humanoAdivina()
	} else {
		maquinaAdivina()
	}
}

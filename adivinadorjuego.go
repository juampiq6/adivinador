package main

import "fmt"

func main() {

	fmt.Print(" ¿ Quiere jugar a pensar(P) o adivinar(a) el número ? : ")
	var juego string
	fmt.Scan(&juego)
	if juego == "a" {
		humanoAdivina()
	} else {
		maquinaAdivina()
	}
}

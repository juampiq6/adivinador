package main

import "fmt"

func main() {

	fmt.Print(" ¿ Quiere jugar a pensar(P) o adivinar(a) el número ? : ") // la P mayuscula indica que es la opcion predeterminada en caso de no ingresar ni "p" ni "a"
	var juego string
	fmt.Scan(&juego)
	if juego == "a" { // si se ingresa "a", el humano adivinará
		humanoAdivina()
	} else { // si se ingresa cualquier otro caracter, la maquina adivinará
		maquinaAdivina()
	}
}

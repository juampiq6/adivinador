package main

import (
	"fmt"
	"strconv"
)

func humanoAdivina() {

	cmbBase := generarCombBase()
	fmt.Print("\n_______________________________________________________________________________________________\n")
	fmt.Print("\n                                    Adivine el número:                                     ")
	fmt.Print("\n_______________________________________________________________________________________________\n")
	for {
		num, _ := preguntarNum()
		var cmb combinacion
		bien, reg := verificarCombinacion(cmb, cmbBase)
		if bien == 4 {
			fmt.Print("\n *** Felicitaciones, adivino el número! *** ")
			break
		} else {
			fmt.Print("BIEN : ", bien, " - REGULAR : ", reg)

		}
	}
}

func verificarCombinacion(cmbIngresada combinacion, cmbBase combinacion) (int, int) {
	var bien int
	var reg int
	for cifra, orden := range cmbIngresada.cifras {
		if value, exists := cmbBase.cifras[cifra]; exists {
			if orden == value {
				bien++
			} else {
				reg++
			}
		}
	}
	return bien, reg
}

func preguntarNum() (string, error) {

	fmt.Print("\nIngrese un número de 4 cifras, que no se repitan, del 0 al 9: ")
	var res string
	fmt.Scan(&res)
	parsed, err := strconv.ParseInt(res, 10, 0)
	return parsed, err
}

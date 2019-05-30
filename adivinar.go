package main

import (
	"fmt"
)

func humanoAdivina() {

	cmbBase := generarCombBase()
	fmt.Print("pss... el número es : " + cmbBase.cifrasToString())
	fmt.Print("\n_______________________________________________________________________________________________\n")
	fmt.Print("\n                                    Adivine el número:                                     ")
	fmt.Print("\n_______________________________________________________________________________________________\n")
	for {
		num, _ := preguntarNum()
		var cmb Combinacion
		cmb.Cifras = ParsearCifras(num)                 // transformamos el string recibido en un mapa de tipo <int:int>, y lo guardamos en el atributo cifras de la combinacion
		bien, reg := VerificarCombinacion(cmb, cmbBase) // comparamos la combinacion ingresada, con la combinacion base, generada al principio
		if bien == 4 {                                  // si la cantidad de bien resultante es 4, el usuario adivino el número
			fmt.Print("\n *** Felicitaciones, adivino el número! *** ")
			break
		} else { // sino, le devuelve la cantidad de bien y regular obtenidos
			fmt.Print("BIEN : ", bien, " - REGULAR : ", reg)

		}
	}
}

func preguntarNum() (string, error) { // preguntamos el numero que el usuario debe ingresar, repetimos hasta que sea valido y lo entregamos como string

	fmt.Print("\nIngrese un número de 4 cifras, que no se repitan, del 0 al 9: ")
	var res string
	fmt.Scan(&res)
	err := ValidarStringNum(res, 4)
	if err != nil {
		fmt.Print("*** ", err, ". Por favor, ingrese solo 4 cifras diferentes, del 0 al 9 ***\n")
		res, err = preguntarNum()
	}
	return res, nil
}

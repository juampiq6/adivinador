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
		var cmb combinacion
		cmb.cifras = parsearCifras(num)
		bien, reg := verificarCombinacion(cmb, cmbBase)
		if bien == 4 {
			fmt.Print("\n *** Felicitaciones, adivino el número! *** ")
			break
		} else {
			fmt.Print("BIEN : ", bien, " - REGULAR : ", reg)

		}
	}
}

func preguntarNum() (string, error) {

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

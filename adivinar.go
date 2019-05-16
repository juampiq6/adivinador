package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func humanoAdivina() {

	cmbBase := generarCombBase()
	fmt.Print("\n_______________________________________________________________________________________________\n")
	fmt.Print("\n                                    Adivine el número:                                     ")
	fmt.Print("\n_______________________________________________________________________________________________\n")
	for {
		num, _ := preguntarNum()
		cmb := parsearCombinacion(num)
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
	err := validarStringNum(res, 4)
	if err != nil {
		fmt.Print("*** ", err, ". Por favor, ingrese solo 4 cifras diferentes, del 0 al 9 ***\n")
		res, err = preguntarNum()
	}
	return res, nil
}

func validarStringNum(str string, cant int) error {
	var err error
	if len(str) != cant {
		err = errors.New("El numero ingresado no es de 4 cifras")
		return err
	}
	_, err = strconv.ParseInt(str, 10, 0)
	if err != nil {
		err = errors.New("Hubo un error transformando los caracteres a numeros")
		return err
	}
	if !elemsUnicos(str) {
		err := errors.New("Las cifras del número no se pueden repetir")
		return err
	}
	return nil
}

func parsearCombinacion(str string) combinacion {
	var comb combinacion
	comb.cifras = make(map[int]int)
	strArr := strings.Split(str, "")
	for i, elem := range strArr {
		res, _ := strconv.ParseInt(elem, 10, 0) //devuelve un int64
		comb.cifras[int(res)] = i               //casteamos a int
	}
	return comb
}

func elemsUnicos(arr string) bool {
	for i, valor := range arr {
		for j, elem := range arr {
			if int(valor) == int(elem) && j != i {
				return false
			}
		}
	}
	return true
}

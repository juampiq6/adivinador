package main

import (
	"fmt"
	"strconv"
)

func maquinaAdivina() {

	fmt.Print("\n_______________________________________________________________________________________________\n")
	fmt.Print("\n                                   Adivinador de número:                                        ")
	fmt.Print("\n_______________________________________________________________________________________________\n")

	var modelo model
	modelo.listCmb = []combinacion{generarCombBase()} // se genera la combinacion inicial, al azar
	for i := 0; true; i++ {
		// lastIndex := len(modelo.listCmb) - 1
		fmt.Print(" ¿ Es su número ? : " + modelo.listCmb[i].cifrasToString())
		fmt.Print("\nCifras BIEN = ")
		fmt.Scan(&modelo.listCmb[i].bien)
		fmt.Print("Cifras REGULAR = ")
		fmt.Scan(&modelo.listCmb[i].regular)

		if modelo.listCmb[i].bien == 4 {
			fmt.Print(" ¡Sabia que lo podia adivinar! ;) ")
			break
		} else {
			modelo.generarNuevaCombinacion()
		}

	}

}

func (m *model) generarNuevaCombinacion() {
	salir := false
	cmb := combinacion{0, 0, siguienteNumeroCifrasUnicas(m.listCmb[len(m.listCmb)-1].cifras)}
	for i := 0; i < 10000; i++ {
		cmb.cifras = siguienteNumeroCifrasUnicas(cmb.cifras)
		for i := range m.listCmb {
			b, r := verificarCombinacion(m.listCmb[i], cmb)
			if b != m.listCmb[i].bien || r != m.listCmb[i].regular {
				salir = false
				break
			} else {
				salir = true
			}
		}
		if salir {
			break
		}
	}
	m.listCmb = append(m.listCmb, cmb)
}

func siguienteNumeroCifrasUnicas(n map[int]int) map[int]int {
	num := 0
	var numstr string
	mult := 1000
	for i := 0; i < 4; i++ {
		num = num + n[i]*mult
		mult = mult / 10
	}
	if num < 999 {
		numstr = strconv.Itoa(num)
	}
	for {
		if num >= 9876 {
			numstr = "0123"
		} else {
			if num < 999 {
				num++
				numstr = strconv.Itoa(num)
				numstr = "0" + numstr
			} else {
				num++
				numstr = strconv.Itoa(num)
			}
		}
		if elemsUnicos(numstr) {
			break
		}
	}
	return parsearCifras(numstr)
}

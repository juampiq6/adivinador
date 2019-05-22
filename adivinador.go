package main

import "fmt"

func maquinaAdivina() {

	fmt.Print("\n_______________________________________________________________________________________________\n")
	fmt.Print("\n                                   Adivinador de número:                                        ")
	fmt.Print("\n_______________________________________________________________________________________________\n")

	var modelo model
	modelo.cifrasPosibles = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	modelo.listCmb[0] = generarCombBase() // se genera la combinacion inicial, al azar

	for i := 0; true; i++ {
		// lastIndex := len(modelo.listCmb) - 1
		fmt.Print(" ¿ Es su número ? : " + modelo.listCmb[i].getCifrasString())
		fmt.Print("Cifras BIEN = ")
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

func (m model) generarNuevaCombinacion() {
	salir := false
	var cmb combinacion
	for {
		cmb = generarCombBase()
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

// func siguienteNumeroCifrasUnicas(n map[int]int) {
// 	for i := range n {
// 		if n[i] == 9 {
// 			if v, exists := n[i-1]; exists && v < 9 {
// 				n[i] = 0
// 				n[i-1]++
// 				elemsUnicos("", n)
// 			} else {

// 			}
// 		}
// 	}

// }

// func (m model) ordenarCombinacion() {}

// func (m model) permutarTodos() {}

// func (m model) popCifraPos(i int) {}

// func (m model) guardarPosIng() {}

// func (m model) guardarPosSal() {}

// func (m model) guardarValIng() {}

// func (m model) guardarValSal() {}

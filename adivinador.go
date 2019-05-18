package main

import "fmt"

func maquinaAdivina() {

	fmt.Print("\n_______________________________________________________________________________________________\n")
	fmt.Print("\n                                   Adivinador de número:                                        ")
	fmt.Print("\n_______________________________________________________________________________________________\n")
	continuar := true
	// var cmbList []combinacion

	var modelo model
	modelo.cifrasPosibles = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	modelo.listCmb[0] = generarCombBase() // se genera la combinacion inicial, al azar

	for i := 0; continuar; i++ {
		lastIndex := len(modelo.listCmb) - 1
		fmt.Print(" ¿ Es su número ? : " + modelo.listCmb[i].getCifrasString() + " --- index : " + string(i))
		fmt.Print("Cifras BIEN = ")
		fmt.Scan(&modelo.listCmb[i].bien)
		fmt.Print("Cifras REGULAR = ")
		fmt.Scan(&modelo.listCmb[i].bien)

		if modelo.listCmb[i].bien == 4 {
			fmt.Print(" ¡Sabia que lo podia adivinar! ;) ")
		} else {
			acc := modelo.verificarRestricciones()
			for i, a := range acc {
				acc[i]()
			}
		}

	}
}

func (m model) verificarRestricciones() map[string]func() {
	var acc map[string]func()
	lastIndex := len(m.listCmb) - 1
	if (m.listCmb[lastIndex].bien + m.listCmb[lastIndex].regular) == 0 {
		return map[string]func(){"permutarTodos": m.permutarTodos}
	} //si no hay ninguna cifra bien ni regular, las cifras no forman parte de la combinacion válida

	if (m.listCmb[lastIndex].bien + m.listCmb[lastIndex].regular) == 4 {
		return map[string]func(){"ordenarCombinacion": m.ordenarCombinacion}
	} //si la suma entre los bien y los regular, nos da 4, quiere decir que hemos encontrado las 4 cifras

	if m.listCmb[lastIndex].bien > m.listCmb[lastIndex-1].bien {
		acc["guardarPosIng"] = m.guardarPosIng
	}

	if m.listCmb[lastIndex].bien < m.listCmb[lastIndex-1].bien {
		acc["guardarPosSal"] = m.guardarPosSal
	}
	//si hay mas bien que la anterior combinacion, como se permuta un solo numero siempre, significa que el numero que agregamos esta en la posicion correcta.      si hay menos bien, el numero que sacamos lo es

	if m.listCmb[lastIndex].regular > m.listCmb[lastIndex-1].regular {
		acc["guardarValIng"] = m.guardarValIng

	}
	if m.listCmb[lastIndex].regular < m.listCmb[lastIndex-1].regular {
		acc["guardarValSal"] = m.guardarValSal
	}
	return acc
} //si hay mas regular que la anterior combinacion, como se permuta un solo numero siempre, significa que el numero que agregamos es de la combinacion correcta.   si hay menos regular, el numero que sacamos lo es

func (m model) ordenarCombinacion() {}

func (m model) permutarTodos() {}

func (m model) popCifraPos(i int) {}

func (m model) guardarPosIng() {}

func (m model) guardarPosSal() {}

func (m model) guardarValIng() {}

func (m model) guardarValSal() {}

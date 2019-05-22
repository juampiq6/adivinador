package main

import "fmt"

func maquinaAdivina() {

	fmt.Print("\n_______________________________________________________________________________________________\n")
	fmt.Print("\n                                   Adivinador de número:                                        ")
	fmt.Print("\n_______________________________________________________________________________________________\n")
	continuar := true
	var modelo model
	modelo.cifrasPosibles = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	modelo.listCmb = []combinacion{generarCombBase()}
	for _, val := range modelo.listCmb[0].cifrasStruct {
		modelo.popCifraPosxValor(val.valor)
	}
	for i := 0; continuar; i++ {
		// lastIndex := len(modelo.listCmb) - 1
		fmt.Print("\n ¿ Es su número ? : " + modelo.listCmb[i].getCifrasString() + "\n")
		fmt.Print("Cifras BIEN = ")
		fmt.Scan(&modelo.listCmb[i].bien)
		fmt.Print("Cifras REGULAR = ")
		fmt.Scan(&modelo.listCmb[i].regular)

		if modelo.listCmb[i].bien == 4 {
			fmt.Print(" ¡Sabia que lo podia adivinar! ;) ")
		} else {
			acc := make(map[string]func())
			if i >= 1 {
				acc = modelo.verificarRestricciones()
				fmt.Print("\naccs = ", acc)
				for i := range acc { //mapeador de string con funcion/es a realizar
					acc[i]()
				}
				fmt.Print("\nCIFRAS POSIBLES : ", modelo.cifrasPosibles)
				fmt.Print("\nPenultima COMB : ", modelo.listCmb[len(modelo.listCmb)-2])
				fmt.Print("\nUltima COMB : ", modelo.listCmb[len(modelo.listCmb)-1], "\n")
			}
			if acc["void"] != nil || acc["ordenarCombinacion"] != nil || acc["permutarTodos"] != nil || i == 0 {
				modelo.permutarCifra(true)
			} else {
				modelo.permutarCifra(false)

			}
		}
	}
}

func (m *model) verificarRestricciones() map[string]func() {
	acc := make(map[string]func())
	lastIndex := len(m.listCmb) - 1
	if (m.listCmb[lastIndex].bien + m.listCmb[lastIndex].regular) == 0 {
		return map[string]func(){"permutarTodos": m.permutarTodos}
	} //si no hay ninguna cifra bien ni regular, las cifras no forman parte de la combinacion válida

	if (m.listCmb[lastIndex].bien + m.listCmb[lastIndex].regular) == 4 {
		return map[string]func(){"ordenarCombinacion": m.ordenarCombinacion}
	} //si la suma entre los bien y los regular, nos da 4, quiere decir que hemos encontrado las 4 cifras

	fmt.Print("\nregulares [", lastIndex-1, "][", lastIndex, "]: ", m.listCmb[lastIndex-1].regular, m.listCmb[lastIndex].regular)
	fmt.Print("\nbien [", lastIndex-1, "][", lastIndex, "]: ", m.listCmb[lastIndex-1].bien, m.listCmb[lastIndex].bien)

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
	if m.listCmb[lastIndex].bien == m.listCmb[lastIndex-1].bien && m.listCmb[lastIndex].regular == m.listCmb[lastIndex-1].regular {
		acc["void"] = m.void
	}

	return acc

} //si hay mas regular que la anterior combinacion, como se permuta un solo numero siempre, significa que el numero que agregamos es de la combinacion correcta.   si hay menos regular, el numero que sacamos lo es

func (m *model) ordenarCombinacion() {}

func (m *model) permutarTodos() {}

func (m *model) guardarPosIng() {
	lastIndex := len(m.listCmb) - 1
	indexUltCambio := m.listCmb[lastIndex].indexCambio
	// m.listCmb[lastIndex].bien++
	m.listCmb[lastIndex].cifrasStruct[indexUltCambio].valorCorrecto = true
	m.listCmb[lastIndex].cifrasStruct[indexUltCambio].posicionCorrecta = true
	fmt.Print("\nGuarda POS [", indexUltCambio, "] : ", m.listCmb[lastIndex].cifrasStruct[indexUltCambio].valor)

}

func (m *model) guardarPosSal() {
	lastIndex := len(m.listCmb) - 1
	indexUltCambio := m.listCmb[lastIndex].indexCambio
	m.listCmb[lastIndex] = m.listCmb[lastIndex-1]
	// m.listCmb[lastIndex+1].cifrasStruct = m.listCmb[lastIndex].cifrasStruct
	// m.listCmb[lastIndex+1].bien = m.listCmb[lastIndex].bien + 1
	// m.listCmb[lastIndex+1].regular = m.listCmb[lastIndex].regular
	// m.listCmb[lastIndex+1].indexCambio = m.listCmb[lastIndex].indexCambio
	m.listCmb[lastIndex].cifrasStruct[indexUltCambio].valorCorrecto = true
	m.listCmb[lastIndex].cifrasStruct[indexUltCambio].posicionCorrecta = true
	fmt.Print("\nVuelve y guarda POS [", indexUltCambio, "] : ", m.listCmb[lastIndex].cifrasStruct[indexUltCambio].valor)

}

func (m *model) guardarValIng() {
	lastIndex := len(m.listCmb) - 1
	indexUltCambio := m.listCmb[lastIndex].indexCambio
	if m.listCmb[lastIndex].cifrasStruct[indexUltCambio].posicionCorrecta {
		for i, cif := range m.listCmb[lastIndex].cifrasStruct {
			if !cif.valorCorrecto {
				m.listCmb[lastIndex].cifrasStruct[i].valor = m.listCmb[lastIndex-1].cifrasStruct[indexUltCambio].valor
				m.listCmb[lastIndex].cifrasStruct[i].valorCorrecto = true
				m.listCmb[lastIndex].indexCambio = i
				fmt.Print("\nGuarda VAL [", i, "] : ", m.listCmb[lastIndex].cifrasStruct[i].valor)
				break
				// m.listCmb[lastIndex].regular++
			}
		}
	} else {
		// m.listCmb = append(m.listCmb, combinacion{m.listCmb[lastIndex].bien, m.listCmb[lastIndex].regular, m.listCmb[lastIndex].cifrasStruct, indexUltCambio})
		m.listCmb[lastIndex].cifrasStruct[indexUltCambio].valorCorrecto = true
		fmt.Print("\nGuarda VAL[", indexUltCambio, "] : ", m.listCmb[lastIndex].cifrasStruct[indexUltCambio].valor)
		// m.pushUltimoIndex()
		// m.listCmb[lastIndex].regular++
	}
}

func (m *model) guardarValSal() {
	lastIndex := len(m.listCmb) - 1
	indexUltCambio := m.listCmb[lastIndex].indexCambio
	if m.listCmb[lastIndex].cifrasStruct[indexUltCambio].posicionCorrecta {
		for i, cif := range m.listCmb[lastIndex].cifrasStruct {
			if !cif.valorCorrecto {
				m.listCmb[lastIndex].cifrasStruct[i].valor = m.listCmb[lastIndex-1].cifrasStruct[indexUltCambio].valor
				m.listCmb[lastIndex].cifrasStruct[i].valorCorrecto = true
				m.listCmb[lastIndex].indexCambio = i
				// m.listCmb[lastIndex].regular++
				fmt.Print("\nVuelve y guarda VAL [", i, "] : ", m.listCmb[lastIndex].cifrasStruct[i].valor)
				break
			}
		}

	} else {
		// m.listCmb = append(m.listCmb, combinacion{m.listCmb[lastIndex].bien, m.listCmb[lastIndex].regular, m.listCmb[lastIndex].cifrasStruct, indexUltCambio})
		m.listCmb[lastIndex].cifrasStruct[indexUltCambio].valor = m.listCmb[lastIndex-1].cifrasStruct[indexUltCambio].valor
		m.listCmb[lastIndex].cifrasStruct[indexUltCambio].valorCorrecto = true
		m.listCmb[lastIndex].regular++
		fmt.Print("\nVuelve y guarda VAL [", indexUltCambio, "] : ", m.listCmb[lastIndex].cifrasStruct[indexUltCambio].valor)

	}

}

func (m *model) permutarCifra(push bool) {
	lastIndex := len(m.listCmb) - 1
	// indexUltCambio := m.listCmb[lastIndex].indexCambio
	m.listCmb = append(m.listCmb, combinacion{m.listCmb[lastIndex].bien, m.listCmb[lastIndex].regular, m.listCmb[lastIndex].cifrasStruct, 0})
	// m.listCmb[lastIndex+1].cifrasStruct = m.listCmb[lastIndex].cifrasStruct
	// m.listCmb[lastIndex+1].bien = m.listCmb[lastIndex].bien
	// m.listCmb[lastIndex+1].regular = m.listCmb[lastIndex].regular
	for {
		r := generarRandom(4)
		if !m.listCmb[lastIndex].cifrasStruct[r].valorCorrecto {
			m.listCmb[lastIndex+1].indexCambio = r
			randInd := generarRandom(len(m.cifrasPosibles) - 1)
			m.listCmb[lastIndex+1].cifrasStruct[r].valor = m.cifrasPosibles[randInd]
			m.popCifPosxI(randInd)
			if push {
				m.cifrasPosibles = append(m.cifrasPosibles, m.listCmb[lastIndex].cifrasStruct[r].valor)
				fmt.Print("pushea ", m.listCmb[lastIndex].cifrasStruct[r].valor, " a valores posibles\n")
			} else {
				fmt.Print("elimina ", m.listCmb[lastIndex].cifrasStruct[r].valor, " de valores posibles\n")
			}
			break
		}
	}
}

func (m *model) popCifPosxI(i int) {
	m.cifrasPosibles = append(m.cifrasPosibles[:i], m.cifrasPosibles[i+1:]...)
}

func (m *model) popCifraPosxValor(val int) {
	for i, cifpos := range m.cifrasPosibles {
		if cifpos == val {
			m.popCifPosxI(i)
		}
	}
}

func (m *model) void() {
	// m.permutarCifra(false)
}

// func (m model) pushUltimoIndex() {
// 	lastIndex := len(m.listCmb) - 1
// 	indexUltCambio := m.listCmb[lastIndex].indexCambio
// 	m.cifrasPosibles = append(m.cifrasPosibles, m.listCmb[lastIndex-1].cifrasStruct[indexUltCambio].valor)
// }

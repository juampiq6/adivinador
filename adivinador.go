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
	modelo.listCmb = []Combinacion{generarCombBase()} // se genera la Combinacion inicial, al azar
	for i := 0; true; i++ {
		fmt.Print(" ¿ Es su número ? : " + modelo.listCmb[i].cifrasToString())
		modelo.preguntarBienRegular(i)
		if modelo.listCmb[i].Bien == 4 {
			fmt.Print(" ¡Sabia que lo podia adivinar! ;) ")
			break
		} else {
			modelo.generarNuevaCombinacion()
		}

	}

}

func (m *model) preguntarBienRegular(i int) { // funcion que preguntar la cantidad de bien y de regular de la combinacion generada.
	for {
		for {
			var bienstr string
			fmt.Print("\nCifras BIEN = ")
			fmt.Scan(&bienstr)
			err := ValidarStringNum(bienstr, 1) // valida que se ingrese un solo numero
			if err != nil {
				fmt.Print("*** ", err, ". Por favor, ingrese un número de una sola cifra del 0 al 4. BIEN + REGULAR = 4 ***\n")
			} else {
				m.listCmb[i].Bien, _ = strconv.Atoi(bienstr) // transforma de string a int
				break
			}
		}
		for {
			var regstr string
			fmt.Print("\nCifras REGULAR = ")
			fmt.Scan(&regstr)
			err := ValidarStringNum(regstr, 1) // valida que se ingrese un solo numero
			if err != nil {
				fmt.Print("*** ", err, ". Por favor, ingrese un número de una sola cifra del 0 al 4. BIEN + REGULAR = 4 ***\n")
			} else {
				m.listCmb[i].Regular, _ = strconv.Atoi(regstr) // transforma de string a int
				break
			}
		}
		if m.listCmb[i].Bien+m.listCmb[i].Regular <= 4 { // valida que bien + regular sea como maximo 4
			break
		}
		fmt.Print("*** Por favor, ingrese un número de una sola cifra del 0 al 4. BIEN + REGULAR = 4 ***\n")
	}
}

func (m *model) generarNuevaCombinacion() { // funcion que genera una nueva combinacion valida (de cifras unicas) y que satisfaga las restricciones de las combinaciones de la lista
	continuar := true
	cmb := Combinacion{0, 0, m.listCmb[len(m.listCmb)-1].Cifras} // crea una combinacion con el numero base
	for continuar {
		cmb.Cifras = siguienteNumeroCifrasUnicas(cmb.Cifras) // le asigna el siguiente numero de cifras unicas
		for i := range m.listCmb {                           // por cada combinacion en la lista
			b, r := VerificarCombinacion(m.listCmb[i], cmb)
			if b != m.listCmb[i].Bien || r != m.listCmb[i].Regular { // si no coincide la cantidad de bien o regular
				continuar = true
				break
			} else { // si coinciden la cantidad de bien y regular
				continuar = false // si pasan todas las combinaciones de la lista, y coinciden la cantidad de bien y regular, saldra la nueva combinacion
			}
		}
	}
	m.listCmb = append(m.listCmb, cmb) // agrega la combinacion a la lista de combinaciones
}

func siguienteNumeroCifrasUnicas(n map[int]int) map[int]int { // funcion que devuelve el siguiente numero, cuyas cifras son unicas
	num := 0
	var numstr string
	mult := 1000
	for i := 0; i < 4; i++ { // transforma el mapa a un int
		num = num + n[i]*mult
		mult = mult / 10
	}
	for {
		if num >= 9876 { // si el numero es el ultimo posible antes de pasar a tener 5 cifras
			numstr = "0123" // volvemos al primero numero posible
		} else {
			if num < 999 { // si el numero tiene menos de 4 cifras
				num++                      // obtenemos el siguiente
				numstr = strconv.Itoa(num) // lo transformamos a string
				numstr = "0" + numstr      // y le agregamos un 0 adelante
			} else { // sino, solo obtenemos el siguiente y lo pasamos a string
				num++
				numstr = strconv.Itoa(num)
			}
		}
		if ElemsUnicos(numstr) { // si ese numero resultante tiene sus cifras unicas
			break
		}
	}
	return ParsearCifras(numstr) // transformar de string a mapa <int:int>
}

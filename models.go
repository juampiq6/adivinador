package main

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type model struct { // el modelo nos sirve para guardar una lista de combinaciones
	listCmb []Combinacion
}

type Combinacion struct { // la combinacion almacena las cifras en un mapa <int:int>, la cantidad de bien y de regular obtenidos por esa combinacion
	Bien    int
	Regular int
	Cifras  map[int]int // el mapa almacena el orden de la cifra en la key, y el valor de la cifra en el value
}

func (c Combinacion) cifrasToString() string { //el objetivo de esta funcion es solo mostrar por pantalla al usuario
	var str string
	for i := 0; i < len(c.Cifras); i++ {
		str = str + strconv.Itoa(c.Cifras[i])
	}
	return str
}

func VerificarCombinacion(cmbIngresada Combinacion, cmbBase Combinacion) (int, int) { //compara dos combinaciones y obtiene la cantidad de bien y de regular
	var bien int
	var reg int
	for orden, cifra := range cmbIngresada.Cifras {
		for o, c := range cmbBase.Cifras {
			if cifra == c { // si los value de los mapas coinciden
				if orden == o { // y si ademas tienen la misma key (orden)
					bien++ // se agrega un bien
					break
				} else { // sino tienen la misma key (orden)
					reg++ // se agrega un regular
					break
				}
			}
		}
	}
	return bien, reg
}

func ParsearCifras(str string) map[int]int { // transforma un string a mapa<int:int>
	cifras := make(map[int]int)      // inicializamos el mapa
	strArr := strings.Split(str, "") //creamos un array de string, conteniendo cada elemento una cifra
	for i, elem := range strArr {
		res, _ := strconv.Atoi(elem) // parseamos de string a int, devuelve un int64
		cifras[i] = int(res)         // casteamos a int, y lo asignamos al mapa, en el orden i
	}
	return cifras
}

func ElemsUnicos(arr string) bool { // verifica que los elementos de un string sean unicos
	for i, valor := range arr {
		for j, elem := range arr {
			if j != i && int(valor) == int(elem) { // si el orden es diferente (para que no verifique el mismo elemento) y el valor es igual, devuleve false
				return false
			}
		}
	}
	return true
}

func generarCombBase() Combinacion { // genera una combinacion aleatoria, de cifras unicas
	var cmb Combinacion
	cmb.Cifras = make(map[int]int) // inicializamos el mapa
	for i := 0; i < 4; i++ {
		cmb.Cifras[i] = generarRandomUnico(10, cmb.Cifras) // asignamos un numero aleatorio del 0 al 10 (no incluido) que no se encuentre en el mapa
	}
	return cmb
}

func generarRandom(hasta int) int { // genera un numero aleatorio del 0 al [hasta] (no incluido) => [0-hasta)
	aleat := rand.New(rand.NewSource(time.Now().UnixNano()))
	return aleat.Intn(hasta)
}

func generarRandomUnico(hasta int, cifras map[int]int) int {
	var r int
	exists := true
	for exists {
		r = generarRandom(hasta)
		for i := 0; i < 4; i++ {
			if cifras[i] == r { // si el numero generado ya se encuentra en el mapa
				exists = true
				break // generara otro random
			} else {
				exists = false // sino cuando i==4, saldra con una cifra random diferente al resto
			}
		}
	}
	return r
}

func ValidarStringNum(str string, cant int) error {
	var err error
	if len(str) != cant { // si la cantidad de caracteres del string no es la deseada
		err = errors.New("El numero ingresado no es de " + strconv.Itoa(cant) + " cifra/s") // crea error
		return err
	}
	_, err = strconv.Atoi(str) // tranforma de string a int
	if err != nil {            // si hubo un error en la tranformacion
		err = errors.New("Hubo un error transformando los caracteres a numeros") // crea error
		return err
	}
	if !ElemsUnicos(str) { // si los elementos del string no son unicos
		err := errors.New("Las cifras del número no se pueden repetir") // crea error
		return err
	}
	return nil // si no encuentra error, devolver error=nil
}

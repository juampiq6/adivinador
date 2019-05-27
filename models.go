package main

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type model struct {
	listCmb []combinacion
}

type combinacion struct {
	bien    int
	regular int
	cifras  map[int]int
}

func (c combinacion) cifrasToString() string {
	var str string
	for i := 0; i < 4; i++ {
		str = str + strconv.Itoa(c.cifras[i])
	}
	return str
}

func verificarCombinacion(cmbIngresada combinacion, cmbBase combinacion) (int, int) {
	var bien int
	var reg int
	for orden, cifra := range cmbIngresada.cifras {
		for o, c := range cmbBase.cifras {
			if cifra == c {
				if orden == o {
					bien++
					break
				} else {
					reg++
					break
				}
			}
		}
	}
	return bien, reg
}

func parsearCifras(str string) map[int]int {
	cifras := make(map[int]int)
	strArr := strings.Split(str, "")
	for i, elem := range strArr {
		res, _ := strconv.ParseInt(elem, 10, 0) //devuelve un int64
		cifras[i] = int(res)                    //casteamos a int
	}
	return cifras
}

func elemsUnicos(arr string) bool {
	for i, valor := range arr {
		for j, elem := range arr {
			if j != i && int(valor) == int(elem) {
				return false
			}
		}
	}
	return true
}

func generarCombBase() combinacion {
	var cmb combinacion
	cmb.cifras = make(map[int]int)
	for i := 0; i < 4; i++ {
		ru := generarRandomUnico(10, cmb.cifras)
		cmb.cifras[i] = ru
	}
	return cmb
}

func generarRandom(hasta int) int {
	aleat := rand.New(rand.NewSource(time.Now().UnixNano()))
	return aleat.Intn(hasta)
}

func generarRandomUnico(hasta int, cifras map[int]int) int {
	var r int
	exists := true
	for exists {
		r = generarRandom(hasta)
		for i := 0; i < 4; i++ {
			if cifras[i] == r {
				exists = true
				break
			} else {
				exists = false
			}
		}
	}
	return r
}

func ValidarStringNum(str string, cant int) error {
	var err error
	if len(str) != cant {
		err = errors.New("El numero ingresado no es de " + strconv.Itoa(cant) + " cifra/s")
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

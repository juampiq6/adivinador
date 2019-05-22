package main

import "strconv"

type model struct {
	listCmb        []combinacion
	cifrasPosibles []int
}

type combinacion struct {
	bien         int
	regular      int
	cifrasStruct [4]cifra
	indexCambio  int
}

func (c combinacion) getCifrasString() string {
	str := ""
	for _, cifra := range c.cifrasStruct {
		str = str + strconv.Itoa(cifra.valor)
	}
	return str
}

type cifra struct {
	valor            int
	posicionCorrecta bool
	valorCorrecto    bool
}

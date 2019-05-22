package main

type model struct {
	listCmb        []combinacion
	cifrasPosibles []int
}

type combinacion struct {
	cifras       map[int]int
	bien         int
	regular      int
	cifrasStruct [4]cifra
	indexCambio  int
}

func (c combinacion) getCifrasString() string {
	var str string
	for _, cifra := range c.cifrasStruct {
		str = str + string(cifra.valor)
	}
	return str
}

type cifra struct {
	valor            int
	posicionCorrecta bool
	valorCorrecto    bool
}

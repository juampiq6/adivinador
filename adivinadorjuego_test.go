package main_test

import (
	"errors"
	"reflect"
	"testing"

	a "github.com/adivinadorjuego"
)

func TestValidarStringNum(t *testing.T) {

	tables := []struct {
		s string
		x int
		r error
	}{
		{"5642", 4, nil},
		{"4456", 4, errors.New("Las cifras del número no se pueden repetir")},
		{"045g", 4, errors.New("Hubo un error transformando los caracteres a numeros")},
		{"4", 1, nil},
		{"012", 4, errors.New("El numero ingresado no es de 4 cifra/s")},
		{"123563", 4, errors.New("El numero ingresado no es de 4 cifra/s")},
	}

	for _, table := range tables {
		res := a.ValidarStringNum(table.s, table.x)
		if res != nil && table.r != nil && res.Error() != table.r.Error() {
			t.Errorf("ValidarStringNum falló con parámetros "+table.s+" , %d . Resulto : %d, deberia ser: %d.", table.x, res.Error(), table.r.Error())
		} else {
			if res != table.r && (res == nil || table.r == nil) {
				t.Errorf("°°°°ValidarStringNum falló con parámetros "+table.s+" , %d . Resulto : %d, deberia ser: %d.", table.x, res, table.r)
			}
		}
	}
}

func TestParsearCifras(t *testing.T) {

	tables := []struct {
		s string
		m map[int]int
	}{
		{"8436", map[int]int{0: 8, 1: 4, 2: 3, 3: 6}},
		{"55555", map[int]int{0: 5, 1: 5, 2: 5, 3: 5, 4: 5}},
		{"12", map[int]int{0: 1, 1: 2}},
		{"", map[int]int{}},
	}

	for _, table := range tables {
		res := a.ParsearCifras(table.s)
		if !reflect.DeepEqual(res, table.m) {
			t.Errorf("°°°°ValidarStringNum falló con parámetros "+table.s+" . Resulto : %d, deberia ser: %d.", res, table.m)
		}
	}
}

func TestElemUnicos(t *testing.T) {

	tables := []struct {
		s string
		b bool
	}{
		{"1235", true},
		{"6684", false},
		{"000", false},
		{"9876543210", true},
		{"12345678901", false},
	}

	for _, table := range tables {
		res := a.ElemsUnicos(table.s)
		if res != table.b {
			t.Errorf("°°°°ValidarStringNum falló con parámetros "+table.s+" . Resulto : %d, deberia ser: %d.", res, table.b)
		}
	}
}

func TestVerificarCombinacion(t *testing.T) {

	tables := []struct {
		cb a.Combinacion
		ci a.Combinacion
		b  int
		r  int
	}{
		{a.Combinacion{0, 0, map[int]int{0: 8, 1: 4, 2: 3, 3: 6}}, a.Combinacion{0, 0, map[int]int{0: 8, 1: 4, 2: 3, 3: 6}}, 4, 0},
		{a.Combinacion{0, 0, map[int]int{0: 1, 1: 2, 2: 3, 3: 4}}, a.Combinacion{0, 0, map[int]int{0: 1, 1: 4, 2: 3, 3: 2}}, 2, 2},
		{a.Combinacion{0, 0, map[int]int{0: 6, 1: 7, 2: 8, 3: 9}}, a.Combinacion{0, 0, map[int]int{0: 2, 1: 3, 2: 4, 3: 0}}, 0, 0},
		{a.Combinacion{0, 0, map[int]int{0: 0, 1: 4, 2: 2, 3: 7}}, a.Combinacion{0, 0, map[int]int{0: 8, 1: 4, 2: 3, 3: 6}}, 1, 0},
		{a.Combinacion{0, 0, map[int]int{0: 8, 1: 4, 2: 3, 3: 6}}, a.Combinacion{0, 0, map[int]int{0: 6, 1: 3, 2: 8, 3: 0}}, 0, 3},
	}

	for _, table := range tables {
		bien, regular := a.VerificarCombinacion(table.cb, table.ci)
		if bien != table.b || regular != table.r {
			t.Errorf("°°°°ValidarStringNum falló con parámetros %d , %d . Resulto : bien = %d , regular = %d , deberia ser: bien = %d , regular = %d.",
				table.cb, table.ci, bien, regular, table.b, table.r)
		}
	}
}

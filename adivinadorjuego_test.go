package main_test

import (
	"errors"
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

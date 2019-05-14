package main

type combinacion struct {
	cifras  map[int]int // map en donde el key es la cifra (0-9) y el value es la posicion (0-3)
	bien    int
	regular int
	// se decidio usar map ya que es mucho más facil buscar valores dentro de la estructura, a comparacion por ejemplo del array, donde tenemos que iterar por cada elemento
}

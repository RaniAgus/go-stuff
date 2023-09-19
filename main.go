package main

import (
	"fmt"

	"github.com/RaniAgus/effectivego/person"
)

func main() {
	agus := person.New("Agus", 24)
	fmt.Println(agus.String())

	var p *[]int = new([]int)      // new reserva memoria y devuelve el puntero
	var v []int = make([]int, 100) // make inicializa un slice y devuelve la estructura

	fmt.Println(p, v)
}

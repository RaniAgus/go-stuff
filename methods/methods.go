package main

import (
	"fmt"
	"math"
)

type Vertice struct {
	X, Y float64
}

// Se pueden agregar métodos a structs
func (v Vertice) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// Para que sean mutables, se pasan por referencia
func (v *Vertice) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// En general, todos los métodos de un struct deberían ser punteros o ninguno
// Esto es más eficiente porque no se copia el struct cada vez que se llama a un método

func main() {
	v := &Vertice{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())
}

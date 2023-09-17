package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ----------------------------------------------

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	// Si tengo un T nil, el método M sigue siendo llamable
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// A pesar de que yo podría hacer v.Abs(), no puedo asignar a = v
	// porque v es un Vertex y no un *Vertex, por lo que no implementa Abser
	// a = v

	fmt.Println(a.Abs())

	// ----------------------------------------------

	// Una variable de tipo interface contiene una referencia a las funciones
	// de un tipo concreto que implementa la interface
	var i I

	i = &T{"Hello"}
	fmt.Printf("(%v, %T)\n", i, i)
	i.M()

	i = F(math.Pi)
	fmt.Printf("(%v, %T)\n", i, i)
	i.M()

	var t *T
	// Acá M() maneja el nil por lo que no explota
	t.M()

	i = t
	// Una interface que contiene un nil, no es nil
	if i == nil {
		fmt.Println("i == nil")
	}
	describe(i)

	t = &T{"Hello"}
	i = t

	// Si una interface contiene un tipo concreto, puedo obtener el valor
	// con una sintaxis similar a un cast
	efe, ok := i.(F)
	if ok {
		fmt.Println("F", efe)
	} else {
		// Si solo hubiera recibido un parámetro, habría explotado
		fmt.Println("Not a F")
	}

	// Puedo verificar si una interface contiene un tipo concreto
	// con una sintaxis similar a un switch
	switch v := i.(type) {
	case *T:
		fmt.Println("T", v)
	case F:
		fmt.Println("F", v)
	}
}

// Una función que recibe una interface vacía puede recibir cualquier cosa
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

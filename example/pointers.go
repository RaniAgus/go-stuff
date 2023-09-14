package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.X, v.Y)

	// También puedo acceder a los campos de un struct con un puntero
	p := &v
	p.X = 10 // Esto es lo mismo que (*p).X = 10!!!!!!!
	fmt.Println(v.X, v.Y)

	// También puedo crear punteros a structs
	q := &Vertex{X: 1, Y: 2}
	fmt.Println(q.X, q.Y)
}

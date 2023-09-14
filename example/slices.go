package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)
	for i := range y {
		y[i] = make([]uint8, dx)
		for j := range y[i] {
			y[i][j] = uint8((i + j) / 2)
		}
	}
	return y
}

type Vx struct {
	Lat, Long float64
}

var m = map[string]Vx{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// Los slices son porciones de un array
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// Si muto un slice, muta el array original
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	pic.Show(Pic)

	// Los maps son mutables
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Para obtener una clave, consulto si está presente con el segundo parametro
	value, isPresent := m["Answer"]
	fmt.Println("The value:", value, "Present?", isPresent)

	// También se pueden crear closures, que son funciones que devuelven funciones
	// con un estado
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, a, b := 0, 0, 1
	return func() int {
		x, a, b = a, b, a+b
		return x
	}
}

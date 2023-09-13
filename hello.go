package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	// Esta es una variable de tipo int. Si no proveo un valor, arranca en 0
	var x int = 5

	// Esta es una variable con el tipo inferido
	y := 7

	fmt.Println("Hello, World!", sum(x, y))

	// esto es un if, no tiene paréntesis
	if y > 6 {
		fmt.Println("y is more than 6")
	} else if x == 5 {
		fmt.Println("x is 5")
	} else {
		fmt.Println("else")
	}

	// Esto es un array de 5 enteros. Es fijo como en C
	var a [5]int
	a[2] = 7

	fmt.Println(a)

	// También se puede inicializar así
	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	// Se pueden crear slices para arrays con tamaño dinámico
	c := []int{5, 4, 3, 2, 1}
	c = append(c, 13) // Retorna un nuevo slice
	fmt.Println(c)

	// Se pueden crear Maps
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")

	fmt.Println(vertices)

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Si quitamos los ; de cada parte del for, se convierte en un while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Se puede iterar sobre un array
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	// Se puede iterar sobre un map
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	// Se puede iterar sobre un map
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}

	// Se pueden crear funciones que retornen más de un valor
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Se pueden crear structs
	p := person{name: "Agus", age: 24}
	fmt.Println(p.name)

	// Se pueden crear punteros
	boke := 6
	// Para luego modificar el valor de la variable igual que en C
	inc(&boke)

	fmt.Println(boke)
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func inc(x *int) {
	*x++
}

// Para correr el programa, ejecutar:
// go run hello.go

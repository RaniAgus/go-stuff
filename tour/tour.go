package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// Like for, the if statement can start with a short statement to execute before the condition.
// Variables declared by the statement are only in scope until the end of the if.

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Print("Go runs on ")

	// en un switch puedo declarar una variable y usarla lol
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Esto es un switch true, es como un if else pero feo
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	deferCount()

	// Esto es algo que ejecuta al final de la función
	defer fmt.Println("world")

	fmt.Println("hello")
}

func deferCount() {
	fmt.Println("counting")

	// Esto es un defer, se ejecuta al final de la función
	// Se pueden apilar, se ejecutan en orden inverso
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

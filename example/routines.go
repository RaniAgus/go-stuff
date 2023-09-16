package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // Debo cerrar un channel solo si el receptor necesita esa info
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// Creo un channel con make
	c := make(chan int)
	// Divido el slice en dos y sumo cada mitad en una goroutine
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// Recibo los resultados de las goroutines
	x, y := <-c, <-c

	// Devuelvo el resultado
	fmt.Println(x, y, x+y)

	// Creo un channel con buffer de 10
	c2 := make(chan int, 10)
	// Lleno el channel con los primeros 10 números de fibonacci
	go fibonacci(cap(c2), c2)
	// Recibo los resultados hasta que el channel se cierre
	for i := range c2 {
		fmt.Println(i)
	}
}

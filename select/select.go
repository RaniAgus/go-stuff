package main

import (
	"fmt"
	"time"
)

func fibo(c, quit chan int) {
	x, y := 0, 1
	for {
		// select es como un switch, pero para channels
		select {
		case c <- x: // Si puedo enviar x al channel c, lo hago
			x, y = y, x+y
		case <-quit: // Si puedo recibir del channel quit, termino
			fmt.Println("quit")
			return
		}
	}
}

func tickTickBoom() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // Recibo del channel 10 números de fibonacci
		}
		quit <- 0 // Envío 0 al channel quit para terminar
	}()
	fibo(c, quit)

	tickTickBoom()
}

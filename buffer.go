package main

import (
	"fmt"
	"time"
)

func main() {
	// Esto es un channel con buffer de 2. Si no lo uso, el channel se bloquea
	// hasta que no haya un receiver
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

	// Voy a crear dos channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Voy a ejecutar dos goroutines que van a enviar un mensaje a cada channel
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	// Ahora voy a recibir los mensajes de los channels
	go func() {
		for {
			// Esto es un select. Recibe el primer channel que tenga un mensaje
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	// Para que no termine el programa, espero a que se presione enter
	fmt.Scanln()
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// Esto es un channel. Es un tipo de dato que se puede compartir entre goroutines
	c := make(chan string)

	// Esto es un wait group. Es un contador de goroutines
	// var wg sync.WaitGroup
	// Esto dice que voy a agregar 1 goroutine al contador
	// wg.Add(1)

	// Si uso go, se ejecuta en un hilo aparte y continua con el hilo principal
	go count("sheep", c)

	// Esto es una función anónima que ejecuto en un hilo aparte
	// go func() {
	// 	count("fish", c)
	// Esto dice que terminé una goroutine
	// 	wg.Done()
	// }()

	// Para que no termine el programa, espero a que se presione enter
	// fmt.Scanln()

	msg := <-c // Esto recibe el valor del channel. Es bloqueante
	fmt.Println(msg)

	// Esto itera sobre el channel hasta que se cierre
	// for {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}

	// 	fmt.Println(msg)
	// }

	// También puedo iterar sobre un channel con range
	for msg := range c {
		fmt.Println(msg)
	}

	// Otra forma de esperar a que terminen las goroutines es con Wait
	// wg.Wait()
}

// Esto es una función que recibe algo y un channel que se compartirá con otra goroutine
func count(algo string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- algo // Esto envía el valor al channel
		time.Sleep(time.Millisecond * 500)
	}

	// Esto finaliza el channel
	close(c)
}

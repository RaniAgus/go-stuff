package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Voy a crear 5 workers, cada uno con su propia cpu
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	// Voy a enviar 40 jobs
	for j := 1; j <= 40; j++ {
		jobs <- j
	}

	// Cierro el channel de jobs
	close(jobs)

	// Recibo los resultados
	for a := 1; a <= 40; a++ {
		fmt.Println(<-results)
	}
}

// Esto es un worker pool. Es un patrón de concurrencia que consiste en tener
// un grupo de goroutines que reciben trabajo a través de un channel y devuelven
// el resultado a través de otro channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		results <- fib(j)
		fmt.Println("worker", id, "finished job", j)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

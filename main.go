package main

import "fmt"

func main() {
	// Este es un bug que va a ser corregido en Go 1.22
	var prints []func()
	for i := 1; i <= 3; i++ {
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, print := range prints {
		print()
	}
}

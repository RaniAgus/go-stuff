package main

import (
	"fmt"
	"strings"
)

// Index es una función que retorna el índice de x en el slice s, o -1 si no está presente.

// El tipo parametrizado T debe ser comparable, es decir, debe ser
// posible comparar dos valores de tipo T con el operador ==.

// comparable es una restricción (constraint) de tipo, que se especifica en la
// lista de parámetros de tipo.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// También puedo crear tipos genéricos, por ejemplo, una lista enlazada
type List[T any] struct {
	Value T
	Next  *List[T]
}

func Push[T any](l **List[T], v T) {
	for *l != nil {
		l = &(*l).Next
	}
	*l = &List[T]{Value: v}
}

func (l *List[T]) String() string {
	sb := strings.Builder{}
	for ; l != nil; l = l.Next {
		sb.WriteString(fmt.Sprintf("%v -> ", l.Value))
	}
	sb.WriteString("nil")
	return sb.String()
}

func main() {
	fmt.Println(Index([]int{1, 2, 3}, 2))                  // 1
	fmt.Println(Index([]string{"hola", "mundo"}, "hola"))  // 0
	fmt.Println(Index([]string{"hola", "mundo"}, "adios")) // -1

	var l *List[int]
	Push(&l, 1)
	Push(&l, 2)
	Push(&l, 3)

	fmt.Println(l) // 1 -> 2 -> 3 -> nil
}

package main

import (
	"fmt"

	"github.com/RaniAgus/effectivego/person"
)

func main() {
	p := person.New("Agus", 24)
	fmt.Println(p.String())
}

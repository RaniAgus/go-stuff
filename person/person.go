package person

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// Como la forma de importar un paquete es con el nombre de la carpeta,
// entonces el nombre de la función debe ser New, no NewPerson
// Esto no genera colisiones y es expresivo porque se utiliza como person.New
func New(name string, age int) *Person {
	return &Person{name, age} // Esto es lo mismo que llamar a new(Person) y luego setear los valores
}

// Los getters no llevan "Get"
func (p *Person) Name() string {
	// En varios lados vamos a encontrar inicializaciones de variables en la misma línea
	// que se hace el if, esto es porque el scope de la variable es solo dentro del if
	// y no queremos que sea global
	if name := p.name; len(name) > 0 {
		return name
	}
	return "No name"
}

func (p *Person) Age() int {
	if p == nil {
		return -1
	} // No hace falta poner else porque ya se retornó en el if
	return p.age
}

// Los setters sí llevan "Set"
func (p *Person) SetAge(age int) {
	if p != nil {
		p.age = age
	}
}

// Las interfaces con un solo método deben llevar el sufijo "-er"
// Esto es para que sea más expresivo, por ejemplo: "Reader", "Writer", "Closer", etc.
type Ager interface {
	Age() int
}

// Si existe una interfaz ya definida que se adapta a lo que necesitamos, entonces
// podemos utilizarla en vez de crear una nueva
// Stringer es una interfaz que tiene un solo método: String()
func (p *Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name(), p.Age())
}

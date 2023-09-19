package person

type Person struct {
	Name string
	Age  int
}

// Como la forma de importar un paquete es con el nombre de la carpeta,
// entonces el nombre de la función debe ser New, no NewPerson
// Esto no genera colisiones y es expresivo porque se utiliza como person.New
func New(name string, age int) Person {
	return Person{name, age}
}

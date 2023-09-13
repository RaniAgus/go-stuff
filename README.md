# go-stuff

- Los proyectos se crean en la carpeta ~/go/src
- Adentro se crea un módulo con `go mod init <nombre_del_paquete>`. Se crea un archivo `mod.go` que se ve así:

```go
module raniagus.github.io/example

go 1.21.1
```

- Este es un hola mundo:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
}
```

- Se puede ejecutar de una sin anestesia usando `go run hello.go`
- Sino, para compilar ejecutamos `go build`, y nos va a crear un `./example`
- `go install` te instala el ejecutable en `~/go/bin`

## Intro

- [Learn Go in 12 Minutes](https://www.youtube.com/watch?v=C8LgvuEBraI)
- [Concurrency in Go](https://www.youtube.com/watch?v=LvgVSSpwND8)

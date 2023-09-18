package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	output := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")

	idx := bytes.Index(input, []byte{x})
	if idx == -1 {
		return x
	}
	return output[idx]
}

func (rot rot13Reader) Read(bytes []byte) (int, error) {
	n, err := rot.r.Read(bytes)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		bytes[i] = rot13(bytes[i])
	}

	return n, nil
}

func main() {
	// io.Reader es una interfaz que representa el stream de datos
	r := strings.NewReader("Hello, Reader!")

	// Creo un buffer que lee de a 8 bytes
	b := make([]byte, 8)
	for {
		// Leo los primeros 8 bytes del stream de datos
		n, err := r.Read(b)

		// Esto devuelve la cantidad de bytes leídos y si hubo un error
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		// Si terminé, finalizo
		if err == io.EOF {
			break
		}
	}

	// Puedo hacer readers que consuman readers
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	rdr := rot13Reader{s}
	io.Copy(os.Stdout, &rdr)
}

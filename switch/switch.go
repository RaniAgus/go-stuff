package switch

// Los switches pueden ser utilizados sin una condición, en este caso se evalúa
// cada case hasta que uno de ellos sea verdadero como un if-else-if-else
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

// Cada case se puede separar con comas:
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

// Los breaks no son necesarios porque el switch termina cuando se encuentra un case
// que sea verdadero
// También, si estamo
func helper(src []byte, validateOnly bool) (err error) {
	Loop:
    for n := 0; n < len(src); n += size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break // Los breaks se pueden utilizar para cortar el switch antes de que termine
            }
            size = 1
            update(src[n])

        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                break Loop // Los breaks también se pueden utilizar para cortar un loop usando labels
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift)
        }
    }
}

// Los switches también pueden ser utilizados para evaluar tipos
// En este caso, el tipo de la variable se evalúa contra el tipo de cada case
// y si coincide, se ejecuta ese case
func do(t interface{}) {
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t)             // t has type bool
	case int:
		fmt.Printf("integer %d\n", t)             // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}

package data

import (
	"bytes"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

// Como la función new inicializa las variables en cero, no hace falta hacerlo
// en el constructor
// - El valor cero de un buffer es un buffer vacío listo para ser utilizado
// - El valor cero de un mutex es un mutex desbloqueado
func New() *SyncedBuffer {
	return new(SyncedBuffer)
}

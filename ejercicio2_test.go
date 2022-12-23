package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromedio(t *testing.T) {

	calificaciones := []float64{6.2, 6.7, 5.6, 5.2, 7.0}

	var esperado float64
	for _, v := range calificaciones {
		esperado += v
	}
	esperado = esperado / float64(len(calificaciones))

	resultado := Promedio(calificaciones...)

	assert.Equal(t, esperado, resultado, "El promedio es diferente")
}

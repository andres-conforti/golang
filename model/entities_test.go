package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComprobarCadena(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
		{"NN0A1111", true, "NN", "000A", 4},
	}

	for _, testData := range cases {
		r, err := ComprobarCadena(testData.Input)
		// acá agregar chequeos propios del test por ejemplo:
		assert.Equal(t, r.Type, testData.Type, "Valores diferentes")
		assert.Equal(t, r.Length, testData.Length, "Valores diferentes")
		assert.Equal(t, r.Value, testData.Value, "Valores diferentes")
		assert.Equal(t, err != nil, testData.Success, "Valores diferentes")
	}
}

//SE EJECUTA CON: go test ./... -count=1 -cover
//Acordate de buscar otros asserts en stackoverflow.

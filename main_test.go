package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//SE EJECUTA CON: go test *.go -count=1
//ESTE ANDA:    go test ./... -count=1
//el count refresca el cache para evitar que usemos un test viejo.

func Testmain(t *testing.T) { /*
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
		}

		for _, testData := range cases {
			r, err := ParseString([]byte(testData.Input))
			// acá agregar chequeos propios del test por ejemplo:
			//assert.Equal(t, err == nil, testData.Success)
		}*/

	var texto string = "TX"
	var numero string = "NN"

	assert.Equal(t, texto, numero, "The two words should be the same.")
}

package main

import (
	"fmt"

	"tudai2021.com/model"
)

func main() {
	//declaro el valor de la cadena a trasnformar en una estructura.
	cadena := "TX00"
	//asigno a la variable "c y err" el llamado a la funcion ComprobarCadena con el valor anterior.
	c, err := model.ComprobarCadena(cadena)
	//si el resultado es correcto, el error seria null (nil aca en lang) (err = nill)
	if err == nil {
		//imprimo el valor de c
		fmt.Println(c)
	} else {
		//en otro caso, imprimo el error utilizando la funcion vista en clase "panic" cortando la funcion del programa.
		panic(err)
	}
}

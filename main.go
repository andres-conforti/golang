package main

import (
	"fmt"
	"regexp"
	"strconv"

	"tpgolang.com/model"
)

//Creo un main donde defino la cadena a usar y la funcion encargada de darmelo.
func main() {
	c := "TX03ABC" //agregar cadena
	comprobarCadena(c)
}

func comprobarCadena(c string) {

	//Comprueba que el LARGO DEL VALOR sea en numeros
	var IsNumber = regexp.MustCompile(`^[0-9]+$`).MatchString
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	a := IsNumber(c[2:4])

	//mientras que la cadena sea mayor o igual a 4, y el largo del valor sea en numeros.
	if len(c) >= 5 && a {
		//Convierto a int, el largo del valor que esta en string.
		length, err := strconv.Atoi(c[2:4])
		fmt.Println(err)
		//mientras que la cadena tenga la misma cantidad de valores que lenght.
		//y sus primeros 2 caracteres sean TX o NN entra
		if len(c[4:]) == length {

			if ((c[:2] == "TX") || (c[:2] == "NN")){
			//creo las siguientes variables para revisar que los valores sean numeros o letras.
			isNN := IsNumber(c[4:])
			isTX := IsLetter(c[4:])

			//si la cadena incia con TX y esta compuesta por solo letras, entra.
			if (c[:2] == "TX") && isTX {
				cadena := model.NewCadena(c[:2], length, c[4:])
				fmt.Println(cadena)
			} else if (c[:2] == "NN") && isNN { 
				//si la cadena incia con NN y esta compuesta por solo numeros, entra.
				cadena := model.NewCadena(c[:2], length, c[4:])
				fmt.Println(cadena)
			} else {
				fmt.Println("NO CORRESPONDE EL TIPO CON SU CONTENIDO")
			}

		} else {
			fmt.Println("NO CORRESPONDE EL TAMAÑO CON LO INDICADO")
		}
	} else {
		fmt.Println("NO ES UNA CADENA VALIDA")
	}
}

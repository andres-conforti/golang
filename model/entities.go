package model

import (
	"errors"
	"regexp"
	"strconv"
)

//estructura de la cadena
type cadena struct {
	Type   string
	Length int
	Value  string
}

//funcion que crea la cadena con los valores que se le envia.
func NewCadena(tipo string, size int, valor string) cadena {
	return cadena{tipo, size, valor}
}

//funcion encargada de comprobar que la cadena sea correcta y devolver una estructura,
// o un error y una estructura vacia en el caso de que no.
func ComprobarCadena(c string) (cadena, error) {
	//declaro la variable lenght como un llamado a la funcion checkNum con el caracter 3y4 de la cadena
	//Al estar encargados de la distancia, se debe revisar que
	//esta funcion checkNum revisa que este compuestos por numeros esos 2 valores.
	//
	stringL := checkNum(c[2:4])
	if !stringL {
		return cadena{}, errors.New("cadena invalida - Lenght incorrecto")
	}
	numL := lenghtToInt(c[2:4])

	//si mi cadena es mayor a 4 (que tenga los primeros 2 valores para el tipo y los otros 2 para la cantidad)
	if len(c) >= 4 {
		//Si coincide el tamaño de mi value con el lenght, Y si el tipo es TX o NN, entra.
		//Cabe aclarar que se tiene en cuenta que el value sea nulo y mi lenght sea 0
		//Nosotros lo tomamos com que era valido eso, ya que al ser 0 de lenght, no deberia haber un value.
		if matchLenght(c, numL) && checkType(c) {
			//asigno la variable tipo y valor con los correspondientes valores de mi cadena.
			tipo := c[0:2]
			valor := c[4:]
			//reviso el tipo y si su valor coincide.
			//en este caso si es del tipo TEXTO y su valor es todo TEXTO(no numeros), entra.
			if tipo == "TX" && checkString(valor) {
				//devuelvo la cadena y que no hubo error.
				return NewCadena(tipo, numL, valor), nil
			} else if tipo == "NN" && checkNum(valor) { //mismo caso que arriba pero con tipo NUMERO
				//devuelvo la cadena y que no hubo error.
				return NewCadena(tipo, numL, valor), nil
			}
		}
		return cadena{}, errors.New("cadena invalida - TIPO o LENGHT invalido")
	}
	return cadena{}, errors.New("cadena invalida")
}

//Devuelve un entero equivalente a cuantos caracteres tendra el Value.
func lenghtToInt(c string) int {
	integer, _ := strconv.Atoi(c)
	return integer
}

//Revisa que el string enviado sea compuesto por caracteres numericos entre el 0 y 9.
//Ademas considera que al ser nil, tambien lo devuelva, en el caso que estemos preguntando value de un lenght 0.
func checkNum(c string) bool {
	IsNumber := regexp.MustCompile(`^[0-9]+$`).MatchString
	if !IsNumber(c) && c != "" {
		return false
	}
	return true
}

//Revisa que el string enviado sea compuesto por caracteres numericos alfabeticos de "a" hasta la "z", ya sea mayuscula o minuscula.
//Ademas considera que al ser nil, tambien lo devuelva, en el caso que estemos preguntando value de un lenght 0.
func checkString(c string) bool {
	IsLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if !IsLetter(c) && c != "" {
		return false
	}
	return true
}

//Revisa que el valor del tipo en la cadena sea "TX" o "NN".
func checkType(c string) bool {
	if (c[0:2] == "TX") || (c[0:2] == "NN") {
		return true
	} else {
		return false
	}
}

//Compara que el value de mi cadena sea igual de largo que el lenght.
//Se tiene en cuenta que el value sea nulo y el lenght 0
func matchLenght(c string, numL int) bool {
	if !(numL == len(c[4:])) {
		return false
	} else if numL == 0 && c[4:] == "" {
		return true
	}
	return true
}

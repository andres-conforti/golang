package model

type cadena struct {
	Type   string
	Length int
	Value  string
}

func NewCadena(tipo string, size int, valor string) cadena {
	return cadena{tipo, size, valor}
}

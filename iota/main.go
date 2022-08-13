package main

import "fmt"

const (
	Online = iota
	Offline
	Maintenance
	Retired
)

const (
	Add = iota
	Substrac
	Multiply
	Divide
)

type Operancion int

func (op Operancion) calcular(izquierda, derecha int) int {

	switch op {
	case Add:
		return izquierda + derecha
	case Substrac:
		return izquierda - derecha
	case Multiply:
		return izquierda * derecha
	case Divide:
		return izquierda / derecha

	}
	panic("operacion no controlada")
}

func main() {
	fmt.Println("sadasdasd")

	add := Operancion(Add)

	fmt.Println("sumar: ", add.calcular(1, 3))
}

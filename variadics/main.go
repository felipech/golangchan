package main

import "fmt"

func multiplicar(numeros ...int) int {

	multiplicacion := 1

	for _, i := range numeros {
		multiplicacion *= i
	}

	return multiplicacion
}

func main() {

	a1 := []int{1, 2, 3}
	a2 := []int{12, 23, 34}

	mergeSlices := append(a1, a2...)

	respuesta := multiplicar(mergeSlices...)

	fmt.Println("variadics ", respuesta)

}

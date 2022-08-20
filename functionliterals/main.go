package main

import (
	"fmt"
	"unicode"
)

func reporte(texto string, info func(i rune) string) {

	for _, i := range texto {
		fmt.Println("letra : ", info(i))
	}
}

func main() {

	info := func(i rune) string {
		if unicode.IsDigit(i) {
			return "digito"
		} else if unicode.IsLetter(i) {
			return "letra"
		} else if unicode.IsPunct(i) {
			return "puntuacion"
		} else if unicode.IsSpace(i) {
			return "spacios"
		} else if unicode.IsControl(i) {
			return "control"
		} else if unicode.IsMark(i) {
			return "mark"
		} else if unicode.IsSymbol(i) {
			return "Symbol"
		} else {
			return "no esta clasificado"
		}
	}

	reporte("primeroo!! 2 2 +", info)
}

package worker

import (
	"bufio"
	"os"
	"strings"
)

type Resultado struct {
	Linea       string
	NumeroLinea int
	Path        string
}

type Resultados struct {
	Inner []Resultado
}

func NuevoResultado(linea string, numeroLinea int, path string) Resultado {
	return Resultado{linea, numeroLinea, path}
}

func FindInFile(path string, find string) *Resultados {
	file, err := os.Open(path)
	if err != nil {
		panic("error al abrir el archivo")
		return nil
	}

	resultados := Resultados{make([]Resultado, 0)}
	scanner := bufio.NewScanner(file)
	numeroLinea := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), find) {
			r := NuevoResultado(scanner.Text(), numeroLinea, path)
			resultados.Inner = append(resultados.Inner, r)

		}
		numeroLinea += 1
	}
	if len(resultados.Inner) == 0 {
		return nil
	} else {
		return &resultados
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	count int
	sync.Mutex
}

func getWord(linea string) []string {
	return strings.Split(linea, " ")
}

func countLetters(palabra string) int {
	letras := 0
	for _, i := range palabra {
		if unicode.IsLetter(i) {
			letras += 1
		}
	}
	return letras
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	totalLetras := Count{}

	var wg sync.WaitGroup

	for {
		if scanner.Scan() {
			linea := scanner.Text()
			palabras := getWord(linea)
			for _, palabra := range palabras {
				copia := palabra
				wg.Add(1)
				go func() {
					totalLetras.Lock()
					defer totalLetras.Unlock()
					defer wg.Done()
					sum := countLetters(copia)
					totalLetras.count += sum

				}()
			}
		} else {
			break
		}
	}

	wg.Wait()
	totalLetras.Lock()
	sum := totalLetras.count
	totalLetras.Unlock()
	fmt.Println("total es: ", sum)

}

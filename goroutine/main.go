package main

import (
	"fmt"
	"sync"
)

func sumar(numeros []int, wg *sync.WaitGroup) {

	var total int

	for i := 0; i < len(numeros); i++ {
		total += numeros[i]
	}
	fmt.Println("Total: ", total)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{12, 2, 355, 44, 54, 6}
	arr3 := []int{15, 26, 37, 43, 5, 6}

	arrTotal := [][]int{arr1, arr2, arr3}

	wg.Add(len(arrTotal))
	for i := 0; i < len(arrTotal); i++ {
		go sumar(arrTotal[i], &wg)
	}
	wg.Wait()

}

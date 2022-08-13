package main

import (
	"fmt"
	"sync"
)

func main() {

	/*primer ejercicio**/
	/*	c := gen()
		recibir(c)*/

	// segundo ejercicio
	salir := make(chan int)
	b := gen2(salir)

	recibir2(b, salir)

	/*ejercicio 3*/

	comaOk := make(chan int)

	go func() {
		comaOk <- 42
	}()
	v, ok := <-comaOk
	fmt.Println("valor ", v, " segundo ", ok)

	close(comaOk)

	v, ok = <-comaOk
	fmt.Println("valor ", v, " segundo ", ok)

	/*
		cuarto ejercicio

	*/

	canal4 := make(chan int)

	go func() {

		for i := 0; i < 100; i++ {
			canal4 <- i
		}
		close(canal4)
	}()

	for v := range canal4 {
		fmt.Println("Ejercicio 4", v)
	}

	/*quinto*/

	crearGoRoutines(10)

}

func recibir2(b <-chan int, salir chan int) {

	for {
		select {
		case v := <-b:
			fmt.Println(v)
		case <-salir:
			return

		}
	}

}

func recibir(c <-chan int) {

	for i := range c {
		fmt.Println("imprimir valores ", i)
	}

}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func gen2(salir chan<- int) <-chan int {
	//canales bidireccionales necesitan su propia goroutine
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}

func crearGoRoutines(nroGoRoutines int) {
	var wg sync.WaitGroup
	wg.Add(nroGoRoutines)

	for i := 0; i < nroGoRoutines; i++ {
		routi := make(chan int)
		go func() {

			for ii := 0; ii < 10; ii++ {
				routi <- ii
			}
			close(routi)
			wg.Done()
		}()

		recibir(routi)
	}

}

package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

type Rectangulo struct {
	largo int
	ancho int
}

func calcularArea(figura Rectangulo) (int, int) {
	area, perimetro := 0, 0

	area = figura.largo * figura.ancho
	perimetro = (2 * figura.largo) + (2 * figura.ancho)

	return area, perimetro
}

func main() {
	felipe := Passenger{"Felipe", 10, false}

	micro := Bus{felipe}

	fmt.Println("micro ", micro)

	fmt.Println("asdasd", felipe)

	rectangulo1 := Rectangulo{4, 2}

	fig1, fig2 := calcularArea(rectangulo1)

	fmt.Println("area ", fig1, " perimetro ", fig2)

}

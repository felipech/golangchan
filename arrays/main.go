package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {

	for i := 0; i < len(rooms); i++ {
		if rooms[i].cleaned {
			fmt.Println("Nombre habitacion limpia ", rooms[i].name)
		}
	}

}

func main() {

	fmt.Println("arrays")
	fmt.Println("forma de declarar")

	var myArr [2]int

	myArr[0] = 10
	myArr[1] = 22

	myArr2 := [3]int{1, 2, 3}
	fmt.Println(myArr2)

	myArr3 := [...]int{2, 3, 4}
	fmt.Println(myArr3)

	myArr4 := [4]int{4, 5, 6}
	fmt.Println(myArr4)

	//ejercicios

	rooms := [...]Room{
		{name: "office", cleaned: true},
		{name: "ops"},
		{name: "Reception"},
		{name: "Warehouse"},
	}

	checkCleanliness(rooms)

}

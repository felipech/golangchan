package main

import (
	"fmt"
)

const (
	Low = iota
	Medium
	High
)

type PriorityQueue[P comparable, V any] struct {
	items     map[P][]V
	priorites []P
}

func (pq *PriorityQueue[P, V]) Add(priority P, value V) {
	//como es un mapa seteo la prioridad como llave y luego le asigno el slice que viene
	//en el append lo que hago es acceder al slice que tiene la llave "priority" y le agrego lo que venga en el "value" V
	pq.items[priority] = append(pq.items[priority], value)
}

// retorna el item con mayor prioridad
func (pq *PriorityQueue[P, V]) Next() (V, bool) {

	for i := 0; i < len(pq.priorites); i++ {
		//copia el elemento del slice de prioridades
		priority := pq.priorites[i]
		items := pq.items[priority] //slice asociado a la llave
		if len(items) > 0 {
			next := items[0]
			pq.items[priority] = items[1:]
			return next, true
		}
	}
	var empty V
	return empty, false
}

func NewPriorityQueue[P comparable, V any](priorities []P) PriorityQueue[P, V] {
	return PriorityQueue[P, V]{items: make(map[P][]V), priorites: priorities}
}

/**parte 2*/

func main() {
	fmt.Println("Generics")
	queue := NewPriorityQueue[int, string]([]int{High, Medium, Low})

	queue.Add(Low, "L-1")
	queue.Add(High, "H-1")
	queue.Add(Low, "L-2")
	queue.Add(Low, "L-3")

	fmt.Println(queue.Next())
}

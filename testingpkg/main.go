package main

type Queue struct {
	items     []int
	capacidad int
}

func New(capacidad int) Queue {
	return Queue{make([]int, 0), capacidad}
}

func (q *Queue) Append(item int) bool {
	if len(q.items) == q.capacidad {
		return false
	}
	q.items = append(q.items, item)
	return true
}

func (q *Queue) Next() (int, bool) {
	if len(q.items) > 0 {
		next := q.items[0]
		q.items = q.items[1:]
		return next, true

	} else {
		return 0, false
	}
}

func main() {

}

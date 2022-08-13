package main

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		//verificando que el largo del Queue siemrpe sea igual al indice
		if len(q.items) != i {
			t.Errorf("incorrect queue element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("Fallo al querer agregar un item %v a la queue", i)
		}
	}
	if q.Append(4) {
		t.Errorf("no se puede agregar un elemento a una Queue que esta llena")
	}
}

func TestNext(t *testing.T) {

	q := New(4)

	for i := 0; i < 4; i++ {
		q.Append(i)
	}

	for i := 0; i < 4; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Se deberia poder obtener el elemento de la Queue")
		}
		if item != i {
			t.Errorf("Lo items estan en orden incorrecto")
		}
	}

	item, ok := q.Next()

	if ok {
		t.Errorf("no deberia tener ningun item. se obtuvo: %v", item)
	}

}

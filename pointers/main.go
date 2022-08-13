package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("counter ", counter)
}

type bag struct {
	items       Items
	securityTag bool
}

type Items struct {
	orden       string
	securityTag bool
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func changeAction(item *Items) {
	if item.securityTag {
		item.securityTag = false
		return
	}
	item.securityTag = true
	return

}

func checkout(items []*Items) {

	for _, it := range items {
		if it.securityTag {
			it.securityTag = false
		}
	}
}

func main() {
	fmt.Println("Pointers part")

	counter := Counter{}

	hola := "hola"
	nuevamente := "nuevemente"

	fmt.Println(hola, nuevamente)

	replace(&hola, "nu", &counter)
	fmt.Println(hola, nuevamente)

	fmt.Println("Generando items")

	item1 := Items{"23232", true}
	item2 := Items{"13232", true}
	item3 := Items{"33232", true}
	item4 := Items{"43232", true}

	itemsList := []*Items{&item1, &item2, &item3, &item4}

	for _, it := range itemsList {
		fmt.Println("items v1: ", it.securityTag)
	}

	fmt.Println("item 1: ", item1)
	changeAction(&item1)
	checkout(itemsList)
	fmt.Println("item 1: ", item1)
	fmt.Println("items: ", itemsList)
	for _, it := range itemsList {
		fmt.Println("items v2: ", it.securityTag)
	}
}

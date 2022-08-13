package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Library struct {
	books     []Book
	members   []Member
	registros []Registro
}

type Registro struct {
	nombreLibro string
	startTime   string
	endTime     string
}

type Book struct {
	title    string
	cantidad int
}

type Member struct {
	name  string
	books []string
}

func checkBook(library *Library, nombreLibro string) bool {
	fmt.Println("direccion de memoria library: ", library)
	for _, i := range library.books {
		if i.title == nombreLibro {
			fmt.Println("si se encontro el libro")
			if i.cantidad > 1 {
				fmt.Println("verificando cantidad")
				return true
			}
			fmt.Println("no hay suficiente cantidad de libros")
			return false
		}
	}
	return false
}

func checkMember(library *Library, nombreMiembro string) bool {
	fmt.Println("direccion de memoria library: ", library)
	for _, i := range library.members {
		if i.name == nombreMiembro {
			fmt.Println("si se encontro el libro")
			return true
		}
	}
	return false
}

func generarCheckOutBook(library *Library, nombreLibro, nombreMiembro string) {
	registro := Registro{nombreLibro: nombreLibro, startTime: time.Now().Format(formatoCortoFecha)}
	library.registros = append(library.registros, registro)

}

func generarCheckInBook(library *Library, nombreLibro, nombreMiembro string) {

}

func generarCheckout(library *Library, nombreLibro string, nombreMiembro string) {
	fmt.Println("-------------------------------------------------")
	fmt.Println("direccion de memoria library v111111: ", *library)
	fmt.Println("-------------------------------------------------")
	//verificar si es miembro

	if checkMember(library, nombreMiembro) {
		//existe
		if checkBook(library, nombreLibro) {
			//existe el libro y cantidad
			for xx, it := range library.books {
				if it.title == nombreLibro {
					fmt.Println("cantidad : ", it.cantidad)
					library.books[xx].cantidad -= 1
					fmt.Println("cantidad : ", it.cantidad)
					generarCheckOutBook(library, nombreLibro, nombreMiembro)
					fmt.Println(" library --- final  : ", library)
					break
				}
			}
			for ff, it := range library.members {
				if it.name == nombreMiembro {
					library.members[ff].books = append(library.members[ff].books, nombreLibro)
					fmt.Println("direccion de memoria library: ", library)
					break
				}
			}
			return
		}
		//no existe cantidad
		fmt.Println("no existe cantidad")
		return

	}
	fmt.Println("no existe como miembro")
	return
}

const formatoCortoFecha = "02-01-2006"

func main() {
	year, month, day := time.Now().Date()

	fecha := time.Now().Format(formatoCortoFecha)

	fmt.Println("fecha . ", fecha)

	book1 := Book{"vuelo", 2}
	book2 := Book{"politica", 3}
	book3 := Book{"Clima", 1}
	book4 := Book{"Quimica", 2}

	books := []Book{book1, book2, book3, book4}
	var booksMembers []string
	member1 := Member{"felipe", booksMembers}
	member2 := Member{"Coni", booksMembers}
	member3 := Member{"Paz", booksMembers}

	members := []Member{member1, member2, member3}

	var registros []Registro

	library := Library{books, members, registros}

	fmt.Println("datos libreria: ")

	for _, it := range library.books {
		fmt.Println("Libros: ", it)
	}
	for _, it := range library.members {
		fmt.Println("Miembros: ", it)
	}
	fmt.Println("***********************")
	fmt.Println("aca empieza el checkout")
	fmt.Println("***********************")
	generarCheckout(&library, "politica", "Paz")

	fmt.Println("datos libreria: ")

	for _, it := range library.books {
		fmt.Println("Libros: ", it)
	}
	for _, it := range library.members {
		fmt.Println("Miembros: ", it)
	}
	builder := strings.Builder{}

	builder.WriteString(strconv.Itoa(year))
	builder.WriteString("-")
	builder.WriteString(month.String())
	builder.WriteString("-")
	builder.WriteString(strconv.Itoa(day))

	fmt.Println("tiempo final : ", builder.String())
}

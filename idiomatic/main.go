package main

import "fmt"

//dos formas de usar los metodos en las estructuras

type Coordenada struct {
	x int
	y int
}

func cambiarCordenada(x, y int, coord *Coordenada) {
	coord.x += x
	coord.y += y
}

// esto es como en java asocio una clase y su metodo
func (coord *Coordenada) cambiarCordenada(x, y int) {
	coord.x += x
	coord.y += y
}

// Distancia aca estamos trabajando con la copia de la cordenada y retornamos una copia nueva con las modificaciones
// aca puede ocuparse mapas y slices si es más comodo q pasar un puntero ya q esas estructuras son punteros de por si
func (c Coordenada) Distancia(cordenadaOriginal Coordenada) Coordenada {
	return Coordenada{c.x - cordenadaOriginal.x, c.y - cordenadaOriginal.y}
}

type Space struct {
	ocupado bool
}

type Estacionamiento struct {
	space []Space
}

type Player struct {
	vida    int
	energia int
	nombre  string
}

func espacioOcupados(estaciona *Estacionamiento, numeroEstacionamiento int) {

	estaciona.space[numeroEstacionamiento-1].ocupado = true

}

func (est *Estacionamiento) estacionamientoReciver(numeroEstacionamiento int) {
	est.space[numeroEstacionamiento-1].ocupado = true
}
func (est *Estacionamiento) estacionamientoLeave(numeroEstacionamiento int) {
	est.space[numeroEstacionamiento-1].ocupado = false
}

func (player *Player) modifiedHealth(vida int) {
	player.vida += vida

}
func (player *Player) modifiedEnergy(energia int) {
	player.energia += energia

}
func (player *Player) modifiedName(nombre string) {
	player.nombre = nombre
}

//Ejercicios de reciver function

func main() {

	coordenada := Coordenada{5, 4}

	cambiarCordenada(2, 3, &coordenada)

	coordenada.cambiarCordenada(2, 3)

	primera := Coordenada{2, 3}
	segunda := Coordenada{3, 4}

	distancia1 := primera.Distancia(segunda)

	fmt.Println("distancia: ", distancia1)

	/**Ejercicio**/

	estacionamientos := Estacionamiento{space: make([]Space, 10)}
	fmt.Println("estacionamientos iniciales : ", estacionamientos)
	estacionamientos.estacionamientoReciver(1)
	espacioOcupados(&estacionamientos, 2)

	player := Player{}

	fmt.Println("player antes de los datos: ", player)

	player.modifiedHealth(100)
	player.modifiedEnergy(100)
	player.modifiedName("Felipe")

	fmt.Println("player despues de los datos: ", player)

}

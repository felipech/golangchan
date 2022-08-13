package main

import "fmt"

type Preparador interface {
	PreparaPlato()
}

type EnviarMecanico interface {
	EnviarVehiculo()
}

type SeleccionarTipoAUto interface {
	TomarSeleccionar() Tomar
}

type Vehiculo struct {
	ModeloVehiculo string
	Tipo           string
}

const (
	pesado = iota
	liiviano
	normal
)

type Tomar int

type Pollo string
type Ensalada string

type Moto string
type Auto string
type Camion string

func (p Pollo) PreparaPlato() {
	fmt.Println("cocinando pollo")
}

func (e Ensalada) PreparaPlato() {
	fmt.Println("cocinando ensalada")
}

func (a Vehiculo) EnviarVehiculo() {
	if a.Tipo == "auto" {
		fmt.Println("Enviar al mecanico de autos")
	} else if a.Tipo == "moto" {
		fmt.Println("Enviar al mecanico de Motos")
	} else if a.Tipo == "camion" {
		fmt.Println("Enviar al mecanico de camion")
	} else {
		panic("no existe el modelo indicado")
	}
}

func SeleccionadorVehiculos(vehiculos []EnviarMecanico) {
	for i := 0; i < len(vehiculos); i++ {
		fmt.Println("Revisando tipo vehiculo")
		vehiculo := vehiculos[i]
		vehiculo.EnviarVehiculo()
		fmt.Println("finalizando tipo vehiculo")
	}
}

func prepararPlatos(platos []Preparador) {
	fmt.Println("Preparando platos")
	for i := 0; i < len(platos); i++ {
		plato := platos[i]
		fmt.Println("Plato: ", plato)
		plato.PreparaPlato()
	}
	fmt.Println("Terminando")
}

func (m Moto) String() string {
	return fmt.Sprintf("Motocicleta: %v", string(m))
}
func (a Auto) String() string {
	return fmt.Sprintf("Auto: %v", string(a))
}
func (c Camion) String() string {
	return fmt.Sprintf("Camion: %v", string(c))
}

func (m Moto) TomarSeleccionar() Tomar {
	return liiviano
}
func (a Auto) TomarSeleccionar() Tomar {
	return normal
}

func (c Camion) TomarSeleccionar() Tomar {
	return pesado
}

func enviarSeleccion(autos SeleccionarTipoAUto) {
	switch autos.TomarSeleccionar() {
	case liiviano:
		fmt.Printf("enviando %v a seleccion liviano\n", autos)
	case pesado:
		fmt.Printf("enviando %v a seleccion pesado\n", autos)
	case normal:
		fmt.Printf("enviando %v a seleccion normal\n", autos)
	}
}

func main() {
	platos := []Preparador{Pollo("Pollo Asado"), Ensalada("Ensalada Cesar")}
	vehiculos := []EnviarMecanico{
		Vehiculo{
			ModeloVehiculo: "ford",
			Tipo:           "auto",
		},
		Vehiculo{
			ModeloVehiculo: "susuky",
			Tipo:           "camion",
		}}

	auto := Auto("susuky")
	camion := Camion("Mercedez benz")
	moto := Moto("Ducati")

	enviarSeleccion(moto)
	enviarSeleccion(camion)
	enviarSeleccion(auto)

	SeleccionadorVehiculos(vehiculos)
	prepararPlatos(platos)
}

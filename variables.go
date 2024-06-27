package main

import "fmt"

func main() {

	//Variable de Entero
	var edad int = 30
	fmt.Println(edad)

	//Variable Flotante
	var altura float64 = 1.75
	fmt.Println(altura)

	//Variable Cadena de Texto
	var nombre string = "Juan"
	fmt.Println(nombre)

	//Variable Booleano
	var esEstudiante bool = true
	fmt.Println(esEstudiante)

	//Variable Arreglo (Array)
	var numeros [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	//Variable Slice
	var frutas []string = []string{"Manzana,", "Pera,", "Naranja"}
	fmt.Println(frutas)

	//Variable Mapa (map)
	var edades map[string]int = map[string]int{"Alice": 25, "Bob": 30}
	fmt.Println(edades)

	//Variable Estructura
	type Persona struct {
		Nombre string
		Edad   int
	}
	var persona Persona = Persona{Nombre: "Sebastian", Edad: 25}
	fmt.Println(persona)

	//Variable Puntero
	var puntero *int

	//Variable interface
	var i interface{} = "Texto"
}

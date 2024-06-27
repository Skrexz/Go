package main

import "fmt"

func main() {

	//Operador suma (+)
	var a int = 10
	var b int = 5

	var suma int = a + b //Resultado es 15
	fmt.Println(suma)

	//Operador resta (-)
	var resta int = a - b //Resultado es 5
	fmt.Println(resta)

	//Operador Multiplicación (*)
	var multiplicacion int = a * b //Resultado es 50
	fmt.Println(multiplicacion)

	//Operador división (/)
	var division int = a / b //Resultado es 2 (División entera)
	fmt.Println(division)

	//Operador Modulo (%)
	var modulo int = a % b //Resultado es 0 (Resto de la división)
	fmt.Println(modulo)

	//Operadores de Asignación

	//1. Asignación Simple (=)
	var x int
	x = 10
	fmt.Println(x)

	//2. Asignación con Suma (+=)
	x += 5 //Ahora x es 15
	fmt.Println(x)

	//3. Asignación con Resta (-=)
	x -= 3 //Ahora x es 12
	fmt.Println(x)

	//4. Asignación con Multiplicación (*=)
	x *= 2 //Ahora x es 24
	fmt.Println(x)

	//5. Asignación con División (/=)
	x /= 4 //Ahora x es 6
	fmt.Println(x)

	//6. Asignación de Módulo (%=)
	x %= 4 //Ahora x es 2
	fmt.Println(x)

}

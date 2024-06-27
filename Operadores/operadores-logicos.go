package main

import "fmt"

func main() {

	//1. Y lógico (&&)
	x := true
	y := false

	resultadoY := x && y
	fmt.Println(resultadoY) //El resultado es false

	//2. O lógico (||)
	resultadoO := x || y //El resultado es true
	fmt.Println(resultadoO)

	//3. Negación lógica (!)
	resultadoNe := !x //Resultado es false
	fmt.Println(resultadoNe)

}

package main

import "fmt"

func main() {

	//1. Igual a (==)
	a := 10
	b := 5

	resultado := (a == b) //El resultado es falso
	fmt.Println(resultado)

	//2. Distinto de (!=)
	resultadoDi := (a != b)
	fmt.Println(resultadoDi) //El resultado es true

	//3. Mayor que (>)
	resultadoMa := (a > b)
	fmt.Println(resultadoMa) //El resultado es true

	//4. Menor que (<)
	resultadoMe := (a < b)
	fmt.Println(resultadoMe) //El resultado es false

	//5. Mayor o igual que (>=)
	resultadoMaOigu := (a >= b)
	fmt.Println(resultadoMaOigu) //El resultado es true

	//6. Menor o igual que (<=)
	resultadoMeoigu := (a <= b)
	fmt.Println(resultadoMeoigu) //El resultado es false
}

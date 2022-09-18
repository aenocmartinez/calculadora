package main

import (
	"calculadora/domain"
	"fmt"
)

func main() {
	fmt.Println("Calculadora")
	c := domain.Calculadora{}
	var numero1 int = 2
	var numero2 int = 9

	fmt.Println("Suma: ", c.Sumar(numero1, numero2))
	fmt.Println("Resta: ", c.Restar(numero2, numero1))
	fmt.Println("Multiplicar: ", c.Multiplicar(numero2, numero1))
	fmt.Println("\n *********** Division ***********")
	fmt.Println("Entera: ", c.DivisionEntera(numero2, numero1))
	fmt.Println("Residuo: ", c.DivisionResiduo(numero2, numero1))
}

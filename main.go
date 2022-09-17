package main

import (
	"calculadora/domain"
	"fmt"
)

func main() {
	fmt.Println("Calculadora")
	c := domain.Calculadora{}
	fmt.Println("Total: ", c.Sumar(2, 9))
}

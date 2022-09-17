package domain

type Calculadora struct {
}

func (c *Calculadora) Sumar(numero1, numero2 int) int {
	return numero1 + numero2
}

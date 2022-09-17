package domain

type Calculadora struct {
}

func (c *Calculadora) Sumar(numero1, numero2 int) int {
	return numero1 + numero2
}

func (c *Calculadora) Restar(numero1, numero2 int) int {
	if numero1 > numero2 {
		return numero1 - numero2
	}
	return 0
}

func (c *Calculadora) Multiplicar(numero1, numero2 int) int {
	return numero1 * numero2
}

func (c *Calculadora) DivisionEntera(numero1, numero2 int) int {
	return numero1 / numero2
}

package main

import "fmt"

func main() {

	var tempType int
	var tempValue float64
	var f float64
	var c float64
	var k float64

	fmt.Println("Escolha o tipo da temperatura passada:\n1 - Kelvi\n2 - Fahnehei\n3 - Celsiu")
	fmt.Scanf("%d", &tempType)
	fmt.Println("Valor da temperatura:")
	fmt.Scanf("%g", &tempValue)

	if tempType == 1 {
		k = tempValue
		f = ((k - 273) * 9 / 5) - 32
		c = k - 273
	}

	if tempType == 2 {
		f = tempValue
		k = ((f - 32) * 5 / 9) + 273
		c = (f - 32) * 5 / 9
	}

	if tempType == 3 {
		c = tempValue
		f = (c * 9 / 5) + 32
		k = c + 273
	}
	//%T - descobrir tipagem
	fmt.Printf("%g K(%T)\n", k, k)
	fmt.Printf("%g °F(%T)\n", f, f)
	fmt.Printf("%g °C(%T)\n", c, c)
}

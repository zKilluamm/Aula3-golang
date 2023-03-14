package main

import "fmt"

func main() {
	//variaveis

	var num1, num2, num3, num4 int

	fmt.Print("Digite o primeiro número:")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número número:")
	fmt.Scanln(&num2)
	fmt.Print("Digite o terceiro número:")
	fmt.Scanln(&num3)
	fmt.Print("Digite o quarto número:")
	fmt.Scanln(&num4)

	//média aritmética

	var media int = (num1 + num2 + num3 + num4) / 4

	fmt.Println("A média é:", media)

}

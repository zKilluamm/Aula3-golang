package main

import "fmt"

func main() {
	//função para inserir dados
	var base int
	var lado int

	fmt.Print("Digite valor da base do retângulo:")
	fmt.Scanln(&base)
	fmt.Println("Digite o valor do lado do retângulo:")
	fmt.Scanln(&lado)

	//função calculo área

	var area_rt int = base * lado

	fmt.Println(" O resultado da operação é:", area_rt)

}

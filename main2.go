package main

import "fmt"

func main() {
	var comp, altura, largura int

	//dados da caixa

	fmt.Print("Digite o comprimento da caixa:")
	fmt.Scanln(&comp)
	fmt.Print("Digite a altura da caixa:")
	fmt.Scanln(&altura)
	fmt.Print("Digite a largura da caixa:")
	fmt.Scanln(&largura)

	//função para calcular volume

	var volume int = comp * altura * largura

	fmt.Println("O volume da caixa é:", volume)

}

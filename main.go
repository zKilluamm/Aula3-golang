package main

import "fmt"

func main() {

	var name string
	var idade int
	var peso int

	fmt.Println("Qual o seu nome?:")
	fmt.Scan(&name)
	fmt.Println("Qual a sua idade?:")
	fmt.Scan(&idade)
	fmt.Println("Qual o seu peso?:")
	fmt.Scan(&peso)
	println("O seu nome é: ", name, ", A sua idade é: ", idade, " E o seu peso é: ", peso)

}

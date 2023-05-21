package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Print("Escreva uma string e receba a quantidade de vagais: ")
	fmt.Scan(&a)
	resultado := contadordevogais(a)
	fmt.Print(resultado)
}
func contadordevogais(a string) int {

	vagais := "aeiouAEIUO"
	var quantidade int
	for _, v := range vagais {
		quantidade += strings.Count(a, string(v))
	}
	return quantidade
}

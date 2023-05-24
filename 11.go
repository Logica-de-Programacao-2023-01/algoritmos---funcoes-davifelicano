// Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro
// se ele for um número primo e falso caso contrário.
// Caso o número seja negativo, retorne um erro
package main

import "fmt"

func main() {
	var p int
	fmt.Print("Escreva um número ")
	fmt.Scan(&p)
	resultado, err := primo(p)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(resultado)
	}
}
func primo(p int) (bool, error) {
	if p < 0 {
		return false, nil
	}
	if p < 2 {
		return false, nil
	}
	for i := 2; i < p; i++ {
		if p%i == 0 {
			return false, nil
		}

	}
	return true, nil
}

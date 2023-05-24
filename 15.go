//Crie uma função que receba um número inteiro
//e um slice de inteiros como parâmetros e retorne verdadeiro se o número estiver presente no slice e falso caso contrário.
//Caso o slice esteja vazio, retorne um erro.
package main

import (
	"fmt"
)

func main() {
	var j int
	fmt.Print("Informe um inteiro para saber se ele esta contido no slice: ")
	fmt.Scan(&j)
	slice := []int{1, 23, 4, 5, 6, 7, 8}
	resultado, err := recb(j, slice)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(resultado)
	}
}
func recb(j int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, fmt.Errorf("Erro, o slice esta vazio")
	}

	for _, contem := range slice {
		if contem == j {
			return true, nil
		}
	}
	return false, nil
}

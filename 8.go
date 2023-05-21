// Escreva uma função que receba um slice de inteiros como parâmetro e
// retorne um novo slice com apenas os números pares contidos no slice.
// Caso o slice esteja vazio, retorne um erro.
package main

import (
	"errors"
	"fmt"
)

func main() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 33, 4, 3, 22}
	result, err := par(sli)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}
}
func par(sli []int) ([]int, error) {
	nova := []int{}
	if len(sli) == 0 {
		return nil, errors.New("Erro")
	}
	for _, cd := range sli {
		if cd%2 == 0 {
			nova = append(nova, cd)
		}
	}
	return nova, nil
}

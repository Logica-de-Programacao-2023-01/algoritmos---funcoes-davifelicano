// Escreva uma função que receba dois slices de inteiros como parâmetros e
// retorne um novo slice contendo apenas os números presentes nos dois slices.
// Caso um dos slices esteja vazio, retorne um erro
package main

import (
	"errors"
	"fmt"
)

func main() {
	sli1 := []int{1, 2, 3, 4, 5, 9, 8, 10, 22}
	sli2 := []int{10, 11, 33, 44, 22, 5}
	result, err := dois(sli1, sli2)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}
}
func dois(sli1 []int, sli2 []int) ([]int, error) {
	if len(sli1) == 0 || len(sli2) == 0 {
		return nil, errors.New("Erro,algum slice esta vazio")
	}
	contem := []int{}
	for _, j := range sli1 {
		for _, d := range sli2 {
			if j == d {
				contem = append(contem, j)
				
			}
		}
	}
	return contem, nil
}

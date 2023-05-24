// Escreva uma função que receba um slice de inteiros como parâmetro
// e retorne um novo slice com os valores ordenados de forma crescente.
// Caso o slice esteja vazio, retorne um erro.
package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	slice := []int{2, 31, 23, 43, 1, 2, 3}
	resultado, err := cres(slice)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(resultado)
	}
}
func cres(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Erro, o slice esta vazio")
	}

	sort.Slice(slice, func(i int, j int) bool {
		return slice[i] < slice[j]
	})
	return slice, nil
}

//sort.Ints ordena de forma direta o slice

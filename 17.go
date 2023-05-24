// Crie uma função que receba um slice de strings como parâmetro
// E retorne uma nova string com as strings em ordem alfabética.
// Caso o slice esteja vazio, retorne um erro.
package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"eu ", "amo", "jujuba"}
	r, err := or(slice)
	if err != nil {
		fmt.Print(err)

	} else {
		fmt.Print(r)
	}
}
func or(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("Erro,slice vazio")
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice, nil
}

//slice[j], slice[j+1] = slice[j+1], slice[j] ordenar

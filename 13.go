// Crie uma função que receba um número inteiro como parâmetro
// e retorne a soma de todos os seus dígitos.
// Caso o número seja negativo, retorne um erro.
package main

import (
	"errors"
	"fmt"
)

func main() {
	var in int
	fmt.Print("Informe um número: ")
	fmt.Scan(&in)
	result, err := soma(in)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}
}
func soma(in int) (int, error) {
	if in < 0 {
		return 0, errors.New("Erro, o número é negativo")
	}
	s := 0
	for in > 0 {
		j := in % 10
		s += j
		in /= 10
	}
	return s, nil

}

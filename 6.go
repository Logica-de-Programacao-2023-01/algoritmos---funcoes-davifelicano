// Escreva uma função que receba um slice de strings como parâmetro
// e retorne uma string com todas as strings concatenadas e separadas por vírgulas.
// Caso o slice esteja vazio, retorne um erro.
package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"davi", "Ama,Juliana"}
	resultado, err := sepvir(slice)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(resultado)
	}
}

func sepvir(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("Erro")
	}
	return strings.Join(slice, ","), nil
}

// Crie uma função que receba uma string como parâmetro
// e retorne um novo slice com todas as palavras contidas na string.
// Considere que as palavras são separadas por espaços em branco.
// Caso a string seja vazia, retorne um erro.
package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	sli := "ola tudo bem"
	nov, err := sep(sli)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(nov)
	}
}
func sep(sli string) ([]string, error) {
	if len(sli) == 0 {
		return nil, errors.New("Erro,slice vazio")
	}

	nov := strings.Split(sli, " ")
	return nov, nil
}

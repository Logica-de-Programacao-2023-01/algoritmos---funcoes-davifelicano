// Escreva uma função que receba um slice de strings como parâmetro
// e retorne uma string contendo apenas as strings que começam com uma letra maiúscula.
// Caso o slice esteja vazio, retorne um erro.
package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func main() {

	slice := []string{"davi Juliana Feliciano"}
	reposta, err := maiu(slice)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(reposta)
	}
}
func maiu(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Erro, o slice esta vazio")
	}
	var resposta strings.Builder
	for _, l := range slice {
		if unicode.IsUpper([]rune(l)[0]) {
			resposta.WriteString(l)
			resposta.WriteString(" ")
		}
	}
	return resposta.String(), nil
}

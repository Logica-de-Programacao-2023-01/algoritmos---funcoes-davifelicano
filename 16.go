// Escreva uma função que receba uma string como
// parâmetro e retorne uma nova string com todas as vogais substituídas por "*".
// Caso a string seja vazia, retorne um erro.
package main

import (
	"fmt"

	"strings"
)

func main() {
	slice := []string{"O rato roeu a roupa do rei de roma"}
	r, err := vg(slice)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(r)
	}

}
func vg(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("Erro,o slice esta vazio")
	}
	r := make([]string, len(slice))
	vogal := "AEIUOaeiou"
	for i, p := range slice {
		for _, j := range vogal {
			p = strings.ReplaceAll(p, string(j), "*")
		}
		r[i] = p
	}
	return r, nil
}

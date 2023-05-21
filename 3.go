// Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.
package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"ola", "batata", "200g de grango", "250g de Arroz"}
	tudin := juntar(a)
	fmt.Print(tudin)
}
func juntar(a []string) string {
	return strings.Join(a, "")
}

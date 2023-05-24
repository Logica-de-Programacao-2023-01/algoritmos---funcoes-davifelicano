package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 12345
	soma, err := somaDigitos(numero)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(soma)
	}
}

func somaDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("O número é negativo")
	}

	soma := 0

	for numero > 0 {
		digito := numero % 10
		soma += digito
		numero /= 10
	}

	return soma, nil
}

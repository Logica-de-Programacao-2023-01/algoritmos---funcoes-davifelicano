// Crie uma função que receba um slice de inteiros e
// um valor inteiro, e retorne a posição do primeiro elemento igual ao valor no slice.
// Caso não encontre, retorne -1.
package main

import "fmt"

func main() {
	var x int
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Print(slice)
	fmt.Print("Escreva um número e tenha de retorno a posição dele : ")
	fmt.Scan(&x)
	resultado := inteiro(slice, x)
	fmt.Print(resultado)
}
func inteiro(slice []int, x int) (int) {
	for posi, n := range slice {
		if n == x {
			return posi
		}
	}
	return -1
}

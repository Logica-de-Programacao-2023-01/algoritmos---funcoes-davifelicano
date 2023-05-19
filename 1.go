// Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.
package main

import "fmt"

func main() {
	slice := []int{2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print(media(slice))
}
func media(slice []int) int {
	soma := 0
	for _, x := range slice {
		soma += x
	}
	return soma / len(slice)
}

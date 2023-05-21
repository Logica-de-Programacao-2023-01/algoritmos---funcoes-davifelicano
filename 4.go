// Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.
package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Print(segundo(sli))
}
func segundo(sli []int) int {
	maior := sli[0]
	segundo := sli[1]
	for j := 2; j < len(sli); j++ {
		if sli[j] > maior {
			segundo = maior
			maior = sli[j]
		} else if sli[j] > segundo && sli[j] < maior {
			segundo = sli[j]
		}
	}

	if segundo > maior {
		maior, segundo = segundo, maior
	}
	return segundo
}

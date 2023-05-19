package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Digite um nome")
	fmt.Scan(&a)
	fmt.Print("Digite um nome")
	fmt.Scan(&b)
	resultado, err := divisao(a, b)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Print(resultado)
}
func divisao(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Não pode")
	}
	return a / b, nil
}

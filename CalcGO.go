package main

import (
	"fmt"
	"os"
)

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Erro: Divisão por zero não é permitida")
	}
	return float64(a) / float64(b), nil

}

func main() {

	var (
		opcao      int
		num1, num2 int
	)

	for {
		fmt.Println("\n~~~ Bem-vindo à Calculadora Simples em Go! ~~~")
		fmt.Print("\n ===z= MENU DE OPÇÕES ====\n[1] - Soma\n[2] - Subtração\n[3] - Multiplicação\n[4] - Divisão\n[0] - Sair\n")

		fmt.Print("\nEscolha qual opção deseja: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&num2)

			fmt.Printf("Resultado: %d\n", sum(num1, num2))

		case 2:
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&num2)

			fmt.Printf("Resultado: %d\n", sub(num1, num2))

		case 3:
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&num2)

			fmt.Printf("Resultado: %d\n", mult(num1, num2))

		case 4:
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&num2)

			resultado, err := div(num1, num2)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Resultado da Divisão: %.2f", resultado)
			}

		case 0:
			fmt.Println("Obrigado por usar nossa calculadora!!")
			os.Exit(0) // Saída normal

		default:
			fmt.Println("Opção inválida!")
			continue

		}
	}

}

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
		fmt.Print("\n ==== MENU DE OPÇÕES ====\n[1] - Soma\n[2] - Subtração\n[3] - Multiplicação\n[4] - Divisão\n[0] - Sair\n")

		fmt.Print("\nEscolha qual opção deseja: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:

			fmt.Printf("\nVocê escolheu SOMA")
			fmt.Print("\nDigite o primeiro número: ")
			fmt.Scanln(&num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&num2)

			resultado := sum(num1, num2)
			fmt.Printf("Resultado da soma: %d\n", resultado)

		case 2:

			fmt.Printf("\nVocê escolheu SUBTRAÇÃO")
			fmt.Print("\nDigite o primeiro número: ")
			fmt.Scanln(&num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&num2)

			resultado := sub(num1, num2)
			fmt.Printf("Resultado da subtração: %d\n", resultado)

		case 3:

			fmt.Printf("\nVocê escolheu MULTIPLICAÇÃO")
			fmt.Print("\nDigite o primeiro número: ")
			fmt.Scanln(&num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&num2)

			resultado := mult(num1, num2)
			fmt.Printf("Resultado da multiplicação: %d\n", resultado)

		case 4:

			fmt.Printf("\nVocê escolheu DIVISÃO")
			fmt.Print("\nDigite o primeiro número: ")
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

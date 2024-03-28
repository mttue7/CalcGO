package main

import (
	"fmt"
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

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Erro: Divisão por zero não é permitido")
	} else {
		return a / b, nil
	}
}

func main() {

	var (
		opcao      int
		num1, num2 int
	)

	fmt.Println("\n~~~ Bem-vindo à Calculadora Simples em Go! ~~~")
	fmt.Print("\n ==== MENU DE OPÇÕES ====\n[1] - Soma\n[2] - Subtração\n[3] - Multiplicação\n[4] - Divisão\n[q] - Sair\n")

	fmt.Print("Escolha qual opção deseja: ")
	fmt.Scanf("%d", &opcao)

	for num1 || num2 !='n'  {
		switch opcao {

		case 1:
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanf("%d\n", &num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanf("%d\n", &num2)
			fmt.Printf("Resultado: %d\n", sum(num1, num2))

		case 2:
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanf("%d\n", &num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanf("%d\n", &num2)
			fmt.Printf("Resultado: %d\n", sub(num1, num2))

		case 3:
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanf("%d\n", &num1)

			fmt.Print("Digite o segundo número: ")
			fmt.Scanf("%d\n", &num2)
			fmt.Printf("Resultado: %d\n", mult(num1, num2))

		case 4:
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanf("%d", &num1)
		
			fmt.Print("Digite o segundo número: ")
			fmt.Scanf("%d", &num2)
		
			resultado, err := div(num1, num2)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Resultado da Divisão:", resultado)
			}
		
		case 'q':
			fmt.Println("\nObrigado por usar nossa calculadora.")
			return

		default:
			fmt.Println("Opção inválida!")

		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(operacion)

	// Separar la operacion en valores
	valores := strings.Split(operacion, "+")
	fmt.Println(valores)
	// Realizar la operacion _ sin conversion queda como texto
	fmt.Println(valores[0] + valores[1])
	// Realizar la operacion _ CON conversion YA son datos numericos
	operador1, err := strconv.Atoi(valores[0])
	if err != nil {
		fmt.Println("Error de conversion para ", valores[0])
	}
	operador2, err := strconv.Atoi(valores[1])
	if err != nil {
		fmt.Println("Error de conversion para ", valores[1])
	}

	fmt.Println("Resultado: ", operador1+operador2)
}

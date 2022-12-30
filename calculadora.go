package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (Calc) Operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		return (operador1 + operador2)
	case "-":
		return (operador1 - operador2)
	case "*":
		return (operador1 * operador2)
	case "/":
		return (operador1 / operador2)
	default:
		fmt.Println(operador, "No esta soportado")
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func LeerEntraada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entrada := LeerEntraada()
	operador := LeerEntraada()
	fmt.Println(entrada)
	fmt.Println(operador)
	c := Calc{}
	fmt.Println(c.Operate(entrada, operador))
}

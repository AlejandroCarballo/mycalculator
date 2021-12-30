package mycalculator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	operationRegexp string = `[0-9]+(\.[0-9]+)?[\+\-\*/][0-9]+(\.[0-9]+)?` //Comprueba que inserte un float operador float
	operatorsRegexp string = `[\+\-\*/]`                                   //Comprueba si se trata de un operador permitido.
)

func GetInput() string {
	fmt.Println("Inserte una operacion")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func ObtainOperation() (string, error) {
	operation := GetInput()
	expr := operationRegexp
	pattern := regexp.MustCompile(expr)
	if pattern.MatchString(operation) {
		//Comprobamos si la entrada concuerda con nuestra expresion, en caso contrario devolvemos un error.
		return operation, nil
	}
	return "", errors.New("Lexical: Input is not an operation")
}

func ObtainNumbers(operation string) (float64, float64) {
	expr := operatorsRegexp
	patern := regexp.MustCompile(expr)
	numbers := patern.Split(operation, 2)
	operand1, _ := strconv.ParseFloat(numbers[0], 32)
	operand2, _ := strconv.ParseFloat(numbers[1], 32)
	return operand1, operand2
}

func ObtainOperator(operation string) byte {
	expr := operatorsRegexp
	patern := regexp.MustCompile(expr)
	indexes := patern.FindIndex([]byte(operation))[0]
	return operation[indexes]
}

func Operate() float64 {
	operation, err := ObtainOperation() //Obtener la operacion del usuario
	if err != nil {
		os.Exit(1)
	}
	value1, value2 := ObtainNumbers(operation)                  //Obtener los valores
	return Calculate(value1, value2, ObtainOperator(operation)) //Devolver el calculo obtenido a partir de los valores y el operador.
}

func Calculate(value1 float64, value2 float64, operator byte) float64 {
	var ans float64
	switch operator {
	case '+':
		ans = value1 + value2
	case '-':
		ans = value1 - value2
	case '*':
		ans = value1 * value2
	case '/':
		ans = value1 / value2
	}
	return ans
}

func main() {
	ans := Operate() //Funcion que realizara el calculo
	fmt.Printf("The output value is %f\n", ans)
}

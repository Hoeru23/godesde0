package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TablaMultiplicar() string {
	fmt.Println("Ingrese el número a generar la tabla de multiplicar")
	for {
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El no. ingresado es incorrecto. Intente nuevamente.")
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}

	return texto
}

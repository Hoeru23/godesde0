package main

import (
	"fmt"

	"github.com/Hoeru23/godesde0/ejercicios"
)

func main() {
	//estado, texto := variables.ConviertoaTexto(1588)
	//fmt.Println(estado)
	//fmt.Println(texto)
	/*
		if os := runtime.GOOS; os == "linux" || os == "OS X." {
			fmt.Println("Esto no es Windows, es ", os)
		} else {
			fmt.Println("Esto es Windows")
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/

	numero, texto, errores := ejercicios.ConversionCondicional("101")
	fmt.Println(numero)
	fmt.Println(texto)
	fmt.Println(errores)
}

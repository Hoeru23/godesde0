package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Hoeru23/godesde0/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo:\n" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, text string) bool {
	arch, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append\n" + err.Error())
		return false
	}

	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Error durante el WriteString\n" + err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error durante el Open\n" + err.Error())
		return
	}

	scanner := bufio.NewScanner((archivo))
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}

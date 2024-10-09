package main

import (
	"github.com/Hoeru23/godesde0/middleware"
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

			numero, texto, errores := ejercicios.ConversionCondicional("101")
			fmt.Println(numero)
			fmt.Println(texto)
			fmt.Println(errores)

		teclado.IngresoNumeros()
	*/

	//iteraciones.Iterar()

	//fmt.Println(ejercicios.TablaMultiplicar())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()

	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencia(2)

	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlice()
	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuario()
	/*
		Pedro := new(m.Hombre)
		e.HumanosRespirando(Pedro)

		Maria := new(m.Mujer)
		e.HumanosRespirando(Maria)
	*/

	//defer_panic.VemosDefer()
	//defer_panic.EjemploPanic()

	/*
		canal1 := make(chan bool)
		go goroutines.MiNombreLentooo("Joel Matos", canal1)
		defer func() {
			<-canal1
		}()
		fmt.Println("Estoy aqui")
	*/

	//webserver.MiWebServer()

	middleware.MiMiddleware()
}

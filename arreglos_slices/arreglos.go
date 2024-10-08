package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}
var matriz [6][6]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54

	tabla2 := [10]string{"Uno", "Dos"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[3][3] = 15
	fmt.Println(matriz)
}

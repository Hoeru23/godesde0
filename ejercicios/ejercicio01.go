package ejercicios

import (
	"strconv"
)

func ConversionCondicional(parametro string) (int, string, error) {
	convertido, errores := strconv.Atoi(parametro)
	var texto string
	if convertido > 100 {
		texto = "Es mayor a 100"
	} else {
		texto = "Es menor a 100"
	}

	return convertido, texto, errores
}

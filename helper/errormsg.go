package helper

import (
	"errors"
	"fmt"
)

/**
const : no valido por el momento
1) Errores de BD
	Error al querer registrar XXXXXXX en la BD
	Error al querer obtener el ID de XXXXXX
	Error al querer obtener las filas afectadas en la tabla XXXXXX
2) Errores de Logica de Negocio
	No valido, el tamaño maximo es de 50 caracteres
	No valido, no se puede usar caracteres especiales
	No valido, No concide con el formato esperado 951000000
**/

/*ErrorCustom modelo general para entregar los errores*/
func ErrorCustom(errTitle string, message string) error {
	return errors.New(fmt.Sprint(errTitle, " # ", message))
	// return fmt.Errorf(errTitle, " # ", message)
}

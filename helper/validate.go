package helper

import (
	"regexp"
	"strings"
)

/*ValidateEmailStr permite validar el Email de tipo string, un tipo hector.limauya@upeu.edu.pe */
func ValidateEmailStr(email string) bool {
	var re = regexp.MustCompile(`^(?:[a-zA-Z0-9])([-_0-9a-zA-Z]+(\.[-_0-9a-zA-Z]+)*|^\"([\001-\010\013\014\016-\037!#-\[\]-\177]|\\[\001-011\013\014\016-\177])*\")@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}\.?$`)
	if len(re.FindStringIndex(email)) > 0 {
		return true
	}
	return false
}

/*ValidatePasswordStr permite validar el password de tipo str, valida caracteres mayusculas y minusculas números y caracteres especiales*/
func ValidatePasswordStr(str string) bool {
	var re = regexp.MustCompile(`[^A-Za-z0-9]`)
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}

/*ValidateStrNumeric permite validar un string con informacion numerica */
func ValidateStrNumeric(str string) bool { //
	var re = regexp.MustCompile(`^[0-9]+$`)
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}

/*ValidateDateStr permite validar la fecha de tipo string, solo almite los siguientes formatos de fechas YYYY/MM/DD o YYYY-MM-DD*/
func ValidateDateStr(str string) bool {
	var re = regexp.MustCompile(`^\d{4}[\-\/\s]?((((0[13578])|(1[02]))[\-\/\s]?(([0-2][0-9])|(3[01])))|(((0[469])|(11))[\-\/\s]?(([0-2][0-9])|(30)))|(02[\-\/\s]?[0-2][0-9]))$`)
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}

/*ValidateTextAndNumberStr valida numeros y letras de un string, no acepta  caracteres especiales*/
func ValidateTextAndNumberStr(str string) bool {
	var re = regexp.MustCompile(`^([A-Za-zÑñáéíóúÁÉÍÓÚ ]+)$`)
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}

/*ValidateTextStr valida caracteres no numericos*/
func ValidateTextStr(str string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z]+$`)
	if len(re.FindStringIndex(str)) > 0 {
		return true
	}
	return false
}

/*ValidateTextComposite permite validar los textos no compuestos, compuestos si es compuesto returna true*/
func ValidateTextComposite(str string) bool {
	strVector := strings.Split(str, " ")
	if len(strVector) > 1 {
		return true
	}
	return false
}

/*FormatFirstUpperStr Limpiar y da formato al texto ejemplo: "Piter Olas" sin UTF-8*/
func FormatFirstUpperStr(str string) string {
	str = strings.Trim(str, " ")
	if len(str) <= 0 {
		return ""
	}
	strVector := strings.Split(strings.ToLower(str), " ")
	str = ""
	for i := len(strVector) - 1; i >= 0; i-- {
		if len(strVector[i]) > 0 {
			strVector[i] = strings.Trim(strVector[i], " ")
			strVector[i] = strings.ToUpper(strVector[i][:1]) + strVector[i][1:len(strVector[i])]
			str = strVector[i] + " " + str
		}
	}
	return strings.Trim(str, " ")
}

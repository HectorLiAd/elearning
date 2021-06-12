package helper

import (
	"fmt"
	"strconv"
	"time"
)

/*ConvStrToDate convierte string a date*/
func ConvStrToDate(date string) (time.Time, error) {
	myDate, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return time.Now(), ErrorCustom(err.Error(), "Error al convertir la fecha")
	}
	return myDate, nil
}

/*RoundDecimal redondea un numero decimal*/
func RoundDecimal(number float64) float64 {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", number), 64)
	fmt.Println(num)
	return num
}

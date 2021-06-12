package helper

import "golang.org/x/crypto/bcrypt"

/*EncryptStr permite cifrar el pwd*/
func EncryptStr(str string, costoEncrypt int) (string, error) {
	// costo := 8 //La eficiencia se relaciona con el costo pero toma tiempo y demora desencriptar
	byte, err := bcrypt.GenerateFromPassword([]byte(str), costoEncrypt)
	return string(byte), err
}

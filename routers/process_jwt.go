package routers

import (
	"strings"

	"github.com/HectorLiAd/elearning/helper"
	"github.com/HectorLiAd/elearning/models"
	"github.com/dgrijalva/jwt-go"
)

/*ProcessJWT Proceso token para extraer sus valores */
func ProcessJWT(tk string) (*models.Claim, bool, string, error) {
	claims, isValidTkn, err := GetClaimsTokenValid(tk)
	if isValidTkn {
		return claims, true, "Encontrado", nil
	}
	return nil, false, string(""), err
}

/*GetClaimsTokenValid de uso general valida y obtenien los claims*/
func GetClaimsTokenValid(tokenString string) (*models.Claim, bool, error) {
	splitToken := strings.Split(tokenString, "Bearer")
	if len(splitToken) != 2 {
		return nil, false, helper.ErrorCustom("TokenError", "Formato de token invalido")
	}
	tokenString = strings.TrimSpace(splitToken[1])
	claims := &models.Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(helper.SecretToken), nil
	})
	if err != nil {
		return nil, false, helper.ErrorCustom("TokenError", err.Error())
	}
	if token.Valid {
		return claims, true, nil
	}
	return claims, false, helper.ErrorCustom("TokenError", "El token ingresado no es correcto")
}

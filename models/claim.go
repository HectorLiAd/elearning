package models

import (
	"github.com/HectorLiAd/elearning/tables/usersystem"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Claim Estructra para procesar el JWT*/
// type Claim struct {
// 	Name    string             `json:"Name"`
// 	UserID  int                `json:"_id"`
// 	Rol     usersystem.Role    `json:"Role"`
// 	Company usersystem.Company `json:"Company"`
// 	jwt.StandardClaims
// }

type Claim struct {
	UserName string             `json:"UserName"`
	Password string             `json:"Password"`
	Role     []*usersystem.Role `json:"Role"`
	Email    string             `json:"Email"`
	Image    string             `json:"Image"`
	UserID   int                `json:"_id"`
	jwt.StandardClaims
}

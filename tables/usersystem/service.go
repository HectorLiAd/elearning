package usersystem

import (
	"strings"
	"time"

	"github.com/HectorLiAd/elearning/helper"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/*Service manejador del repositorio persona*/
type Service interface {
	getUserExists(params *loginUserRequest) (*User, bool, error)
	LoginUser(params *loginUserRequest) (interface{}, error)
	generateJWT(user *User, roleUse []*Role) (string, error)
}

type service struct {
	repo Repository
}

/*NewService retornar el servicio*/
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (serv *service) getUserExists(params *loginUserRequest) (*User, bool, error) {
	userExists, err := serv.repo.GetUserByEmail(params.Email)
	if err != nil {
		return nil, false, err
	}
	pwdUserSolidAcces := []byte(params.Password)
	pwdUserDB := []byte(userExists.PasswordUser)
	err = bcrypt.CompareHashAndPassword(pwdUserDB, pwdUserSolidAcces)
	if err != nil {
		return nil, false, helper.ErrorCustom("PasswordError", "El password es incorrecto")
	}
	return userExists, true, nil
}

func (serv *service) LoginUser(params *loginUserRequest) (interface{}, error) {
	//VALIDAR Y LIMPIAR DATOS Y ENTREGAR ERROR PERSONALIZADO
	params.Email = strings.TrimSpace(params.Email)
	if len(params.Email) > 50 {
		return nil, helper.ErrorCustom("EmailError", "No valido, el tamaño máximo es de 50 caracteres")
	}

	//BUSCAR AL USUARIO POR EL EMAIL
	userExistBool, err := serv.repo.ExistUserEmail(params.Email)
	if err != nil {
		return nil, err
	}
	if !userExistBool { //MANDAR EL ERROR PERSONALIZADO
		return nil, helper.ErrorCustom("EmailError", "El email ingresado no existe")
	}
	//VERIFICAR EL USUARIO Y CONTRASEÑA
	user, _, err := serv.getUserExists(params)
	if err != nil {
		return nil, err
	}

	//BUSCAR EL ROL DEL USUARIO
	userRole, err := serv.repo.GetRolUserByID(user.PersonID)
	if err != nil {
		return nil, err
	}

	//GENERAR JWT
	tokenUserStr, err := serv.generateJWT(user, userRole)
	if err != nil {
		return nil, helper.ErrorCustom(err.Error(), "Error al generar el token de autentificación")
	}
	tokenUser := &TokenLogin{
		Token: tokenUserStr,
	}
	return tokenUser, err
}

func (serv *service) generateJWT(user *User, roleUser []*Role) (string, error) {

	//OBTENER DATA PARA PODER PONER A LOS CLAIMS
	miClave := []byte(helper.SecretToken) //Creando clave privada
	claims := jwt.MapClaims{}
	claims["UserName"] = user.UserName
	claims["Password"] = "#Raaaaa!!!"
	claims["Role"] = roleUser
	claims["Email"] = user.Email
	claims["Image"] = user.ImageURL
	claims["_id"] = user.PersonID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims) //El goritmo para encriptar "header"
	tokenStr, err := token.SignedString(miClave)                //Firmando con el string de firma
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

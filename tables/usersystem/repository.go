package usersystem

import (
	"database/sql"

	"github.com/HectorLiAd/elearning/helper"
)

/*Repository nos sirve para hacer las consultas directas con la BD*/
type Repository interface {
	ExistUserEmail(name string) (bool, error)
	GetUserByEmail(name string) (*User, error)
	GetRolUserByID(userID int) ([]*Role, error)
	GetCompanyByID(companyID int) (*Company, error)
}
type repository struct {
	db *sql.DB
}

/*NewRepository crear el nuevo repositorio y retorna con la BD conectada*/
func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{
		db: dataBaseConnection,
	}
}

func (repo *repository) GetUserByEmail(email string) (*User, error) {
	const queryStr = `select persona_id, user_name, email, clave, avatar, estado from usuario where email = ?`
	results := repo.db.QueryRow(queryStr, email)
	user := &User{}
	err := results.Scan(&user.PersonID, &user.UserName, &user.Email, &user.PasswordUser, &user.ImageURL, &user.StateID)
	if err != nil {
		return nil, helper.ErrorCustom(err.Error(), "Error al querer obtener el usuario de la BD")
	}
	return user, err
}

func (repo *repository) ExistUserEmail(email string) (bool, error) {
	const queryStr = `select count(*) from usuario where email = ?`
	result := repo.db.QueryRow(queryStr, email)
	userCant := 0
	err := result.Scan(&userCant)
	if err != nil {
		return false, helper.ErrorCustom(err.Error(), "Error al consultar al usuario en la BD")
	}
	if userCant > 0 {
		return true, nil
	}
	return false, nil
}

func (repo *repository) GetRolUserByID(userID int) ([]*Role, error) {
	const queryStr = `select rol_id, rol_nombre from vw_usuario_rol_modulo_accion where  persona_id = ? group by rol_id`
	results, err := repo.db.Query(queryStr, userID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, helper.ErrorCustom(err.Error(), "Error al querer obtener el detalle de venta de la BD")
		}
	}
	var roles []*Role
	for results.Next() {
		roleUser := &Role{}
		err := results.Scan(&roleUser.RoleID, &roleUser.Name)
		if err != nil {
			return nil, helper.ErrorCustom(err.Error(), "Error al querer obtener el rol del usuario de la BD")
		}
		roles = append(roles, roleUser)
	}

	return roles, nil
}

func (repo *repository) GetCompanyByID(companyID int) (*Company, error) {
	const queryStr = `SELECT ID_SUCURSAL, CODIGO FROM SUCURSAL WHERE ID_SUCURSAL = ?`
	result := repo.db.QueryRow(queryStr, companyID)
	companyUser := &Company{}
	err := result.Scan(&companyUser.CompanyID, &companyUser.Code)
	if err != nil {
		return nil, helper.ErrorCustom(err.Error(), "Error al querer obtener el rol del usuario de la BD")
	}
	return companyUser, nil
}

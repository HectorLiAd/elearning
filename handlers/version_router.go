package handlers

import (
	"database/sql"
	"net/http"

	// modultable "github.com/HectorLiAd/e_learning/tables/modultable"
	// operationmodul "github.com/HectorLiAd/e_learning/tables/operationmodul"
	// person "github.com/HectorLiAd/e_learning/tables/person"
	// "github.com/HectorLiAd/e_learning/tables/role"
	// "github.com/HectorLiAd/e_learning/tables/roleoperation"
	"github.com/HectorLiAd/elearning/tables/usersystem"
	"github.com/go-chi/chi"
)

/*RouterV1 nos permite usar  las rutas del proyecto*/
func RouterV1(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	// r.Use()

	var (
		userSystemRepository = usersystem.NewRepository(db)
		// personRepository         = person.NewRepository(db)
		// roleRepository           = role.NewRepository(db)
		// roleOperationRepository  = roleoperation.NewRepository(db)
		// roleUserRepository       = userrole.NewRepository(db)
		// modulTableRepository     = modultable.NewRepository(db)
		// operationModulRepository = operationmodul.NewRepository(db)
	)
	var (
		userSystemService = usersystem.NewService(userSystemRepository)
		// personService         = person.NewService(personRepository)
		// roleService           = role.NewService(roleRepository)
		// roleOperationService  = roleoperation.NewService(roleOperationRepository)
		// roleUserService       = userrole.NewService(roleUserRepository)
		// modulTableService     = modultable.NewService(modulTableRepository)
		// operationModulService = operationmodul.NewService(operationModulRepository)
	)

	// r.Mount("/person", person.MakeHTTPSHandler(personService, db, moduls.Person()))
	// r.Mount("/role", role.MakeHTTPSHandler(roleService, db, moduls.Role()))
	// r.Mount("/roleOperation", roleoperation.MakeHTTPSHandler(roleOperationService, db, moduls.RoleOperation()))
	// r.Mount("/roleUser", userrole.MakeHTTPSHandler(roleUserService, db, moduls.RoleUser()))
	// r.Mount("/modulTable", modultable.MakeHTTPSHandler(modulTableService, db, moduls.ModulTable()))
	// r.Mount("/operationModul", operationmodul.MakeHTTPSHandler(operationModulService, db, moduls.OperationModul()))

	r.Mount("/accesoSistema", usersystem.MakeHTTPSHandler(userSystemService))

	return r
}

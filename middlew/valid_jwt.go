package middlew

import (
	"net/http"

	"github.com/HectorLiAd/elearning/routers"
	"github.com/HectorLiAd/elearning/tables/usersystem"
)

/*ValidJWT valida el TOKEN*/
func ValidJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessJWT(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*CommonMiddlewareUserRole de otra forma el JWT*/
func CommonMiddlewareUserRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessJWT(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token pipipipipi "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*UseValidRole validar el rol*/
func UseValidRole(next http.Handler, s usersystem.Service) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessJWT(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token pipipipipi "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})

}

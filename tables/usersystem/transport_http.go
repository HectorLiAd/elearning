package usersystem

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler sirve para hacer peticiones de metodos*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	loginUserHandler := kithttp.NewServer(
		makeLoginUserEndPoint(s),
		loginUserRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodPost, "/userSesion", loginUserHandler)

	mensaje := kithttp.NewServer(
		mensaje(s),
		mensajeDec,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/mensaje", mensaje)

	return r
}

func loginUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := loginUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

func mensajeDec(context context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/HectorLiAd/elearning/database"
	verionesrouter "github.com/HectorLiAd/elearning/handlers"
	"github.com/HectorLiAd/elearning/helper"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.InitDB()
	fmt.Println("------Success BD access------")
	defer db.Close()

	r := chi.NewRouter()
	r.Use(helper.GetCors().Handler)
	r.Mount("/v1", verionesrouter.RouterV1(db))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("PUERTO :" + port)
	http.ListenAndServe(":"+port, r)
}

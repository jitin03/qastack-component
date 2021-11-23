package app

import (
	"context"
	"encoding/json"
	"fmt"
	"qastack-components/domain"
	"qastack-components/service"

	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"

	//"log"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	//"os"
)

func getDbClient() *sqlx.DB {
	client, err := sqlx.ConnectContext(context.Background(), "postgres", "host=localhost port=5433 user=postgres dbname=postgres sslmode=disable password=qastack")
	if err != nil {
		panic(err)
	}
	return client
}



func Start() {

	//sanityCheck()

	router := mux.NewRouter()
	dbClient := getDbClient()

	router.Use()
	componentRepositoryDb := domain.NewComponentRepositoryDb(dbClient)
	//wiring
	//u := ComponentHandler{service.NewUserService(userRepositoryDb,domain.GetRolePermissions())}

	c := ComponentHandler{service.NewComponentService(componentRepositoryDb)}

	// define routes

	router.HandleFunc("/api/component/health", func (w http.ResponseWriter,r *http.Request) {
		json.NewEncoder(w).Encode("Running...")
	})

	router.
		HandleFunc("/api/component/add", c.AddComponent).
		Methods(http.MethodPost).Name("AddComponent")

	router.
		HandleFunc("/api/components", c.AllComponent).
		Methods(http.MethodGet).Name("AllComponent")

	router.
		HandleFunc("/api/component/delete/{id}", c.DeleteComponent).
		Methods(http.MethodDelete).Name("DeleteComponent")

	router.
		HandleFunc("/api/component/update/{id}", c.UpdateComponent).
		Methods(http.MethodPut).Name("UpdateComponent")

	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedHeaders: []string{ "Content-Type", "Authorization","Referer"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET","PUT","DELETE","POST"},
	})

	handler := cor.Handler(router)
	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	//logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	if err := http.ListenAndServe(":8091", handler); err != nil {
		fmt.Println("Failed to set up server")

	}

}



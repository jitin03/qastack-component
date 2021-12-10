package app

import (

	"encoding/json"
	"fmt"
	"os"
	"qastack-components/domain"
	logger "qastack-components/loggers"
	"qastack-components/service"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"

	//"log"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	//"os"
)

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	//dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbAddr, 5432, dbUser, dbPasswd, dbName)
	logger.Info(psqlInfo)
	client, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return  client
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
	if err := http.ListenAndServe(":8093", handler); err != nil {
		fmt.Println("Failed to set up server")

	}

}



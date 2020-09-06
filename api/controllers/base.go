package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/skwongg/jackt/api/models"
)

//Server is the instance of the api server
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

//Initialize is the preflight check before the server is started
func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	if DbDriver == "postgres" {
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", DbUser, DbPassword, DbName, DbPort, "", "America/Los_Angeles")
		server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DbDriver)
		}
	} else {
		log.Fatal("\nError, DBDriver not supported.")
	}
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Lift{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

//Run starts the server instance.
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

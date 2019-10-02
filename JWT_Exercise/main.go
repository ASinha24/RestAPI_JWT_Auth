package main

import (
	"GO/JWT/JWT_Exercise/controllers"
	"GO/JWT/JWT_Exercise/driver"
	"GO/JWT/JWT_Exercise/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}
	router := mux.NewRouter()
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	//router.HandleFunc("/protected", controller.TokenVerifyMiddleWare(controller.ProtectedEndpoint())).Methods("GET")
	router.HandleFunc("/protected", utils.TokenVerifyMiddleWare(controller.ProtectedEndpoint())).Methods("GET")

	log.Println("listening to the port :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

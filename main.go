package main

import (
  "log"
	"net/http"
	"erago/controllers"
	"erago/database"
	"erago/entity"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "go",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Product{})
}

func main() {
  initDB()
	log.Println("Starting the HTTP server on port 8080")
	router := mux.NewRouter().StrictSlash(true)
  initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}


func initaliseHandlers(router *mux.Router) {
  router.HandleFunc("/create", controllers.CreateProduct).Methods("POST")
  router.HandleFunc("/get", controllers.GetAllProduct).Methods("GET")
  router.HandleFunc("/getTerbaru", controllers.GetAllProductTerbaru).Methods("GET")
  router.HandleFunc("/getTermurah", controllers.GetAllProductTermurah).Methods("GET")
  router.HandleFunc("/getTermahal", controllers.GetAllProductTermahal).Methods("GET")
  router.HandleFunc("/getAllProductAsc", controllers.GetAllProductAsc).Methods("GET")
  router.HandleFunc("/getAllProductDesc", controllers.GetAllProductDesc).Methods("GET")
  router.HandleFunc("/get/{id}", controllers.GetProductByID).Methods("GET")
  http.Handle("/", router)
}

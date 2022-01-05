package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"erago/database"
	"erago/entity"
	"github.com/gorilla/mux"
)


func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var product []entity.Product
	database.Connector.Find(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var product entity.Product
	database.Connector.First(&product, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}


func CreateProduct(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Product
	json.Unmarshal(requestBody, &product)

	database.Connector.Create(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetAllProductTerbaru(w http.ResponseWriter, r *http.Request) {
	var product []entity.Product
	database.Connector.Order("id asc").Find(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}


func GetAllProductTerlama(w http.ResponseWriter, r *http.Request) {
	var product []entity.Product
	database.Connector.Order("id desc").Find(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func GetAllProductTermurah(w http.ResponseWriter, r *http.Request) {
	var product []entity.Product
	database.Connector.Order("price asc").Find(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
func GetAllProductTermahal(w http.ResponseWriter, r *http.Request) {
	var product []entity.Product
	database.Connector.Order("price desc").Find(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func GetAllProductAsc(w http.ResponseWriter, r *http.Request) {
	var product []entity.Product
	database.Connector.Order("name asc").Find(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
func GetAllProductDesc(w http.ResponseWriter, r *http.Request) {
	var product []entity.Product
	database.Connector.Order("name desc").Find(&product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

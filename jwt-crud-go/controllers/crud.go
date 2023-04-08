package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tentangKode/product"

	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := product.ConnectDBProduct()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "not connected to the database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	var product []product.Product
	if err = db.Find(&product).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	w.Write(response)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := product.ConnectDBProduct()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "not connected to the database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	vars := mux.Vars(r)
	fmt.Println(vars)
	id := vars["id"]

	var product product.Product
	if err = db.First(&product, id).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	} 
		w.Write(response)
	
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := product.ConnectDBProduct()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "not connected to the database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	var product product.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	if err = db.Create(&product).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"message": "data was not successfully added"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	} 
		response := map[string]string{"message": "data added successfully"}
		json.NewEncoder(w).Encode(response)
	
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := product.ConnectDBProduct()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "not connected to the database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var product product.Product
	err = json.NewDecoder(r.Body).Decode(&product)

	if db.Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		response, _ := json.Marshal(map[string]string{"message": "data cannot be changed"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}  
		response := map[string]string{"message": "data changed successfully"}
		json.NewEncoder(w).Encode(response)
	
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	db, err := product.ConnectDBProduct()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "not connected to the database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	var product product.Product
	if db.Delete(&product, "id = ?", id).RowsAffected == 0 {
		response, _ := json.Marshal(map[string]string{"message": "data was not successfully deleted"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}
		response := map[string]string{"message": "data deleted successfully"}
		json.NewEncoder(w).Encode(response)
	

}

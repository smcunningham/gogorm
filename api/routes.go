package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gogorm/models"
	"net/http"
	"strconv"
)

var db *gorm.DB

func init() {
	db = models.InitDB()
}

func Route() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/orders", createOrder).Methods(http.MethodPost)
	r.HandleFunc("/orders/{orderId}", getOrder).Methods(http.MethodGet)
	r.HandleFunc("/orders", getOrders).Methods(http.MethodGet)
	r.HandleFunc("/orders/{orderId}", updateOrder).Methods(http.MethodPut)
	r.HandleFunc("/orders/{orderId}", deleteOrder).Methods(http.MethodDelete)

	return r
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	db.Create(&order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]

	var order models.Order
	db.Preload("Items").First(&order, inputOrderID)
	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []models.Order
	db.Preload("Items").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder models.Order
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	db.Save(&updatedOrder)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedOrder)
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	// Convert `orderId` string param to uint64
	id64, _ := strconv.ParseUint(inputOrderID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("order_id = ?", idToDelete).Delete(&models.Item{})
	db.Where("order_id = ?", idToDelete).Delete(&models.Order{})
	w.WriteHeader(http.StatusNoContent)
}

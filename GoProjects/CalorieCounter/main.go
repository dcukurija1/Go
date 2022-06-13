package main

import (
	"log"
	"fmt"
	"example.com/calorieCounter/food"
	"example.com/calorieCounter/db"
	"example.com/calorieCounter/handlers"
	"net/http"
    "github.com/gorilla/mux"
)

func testFunction(){
	var p food.Plate
	p.AddToPlate(*food.NewFood(1, "Apple", "Fruit", 50, 100))
	p.ShowPlate()
	p.AddToPlate(*food.NewFood(1, "Orange", "Fruit", 60, 100))
	p.ShowPlate()
	fmt.Println("Total calories: ", p.CurrentCalories())
}
func main() {
	DB := db.Init()
    h := handlers.New(DB)
    router := mux.NewRouter()

    router.HandleFunc("/food", h.GetAllFood).Methods(http.MethodGet)
    router.HandleFunc("/food/{id}", h.GetFood).Methods(http.MethodGet)
    router.HandleFunc("/food", h.AddFood).Methods(http.MethodPost)
    router.HandleFunc("/food/{id}", h.UpdateFood).Methods(http.MethodPut)
    router.HandleFunc("/food/{id}", h.DeleteFood).Methods(http.MethodDelete)

    log.Println("API is running!")
    http.ListenAndServe(":4000", router)
}

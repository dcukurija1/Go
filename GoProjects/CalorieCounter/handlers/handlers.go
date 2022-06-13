package handlers

import (
    "gorm.io/gorm"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "example.com/calorieCounter/food"
)

type handler struct {
    DB *gorm.DB
}

func New(db *gorm.DB) handler {
    return handler{db}
}


func (h handler) GetAllFood(w http.ResponseWriter, r *http.Request) {}

func (h handler) GetFood(w http.ResponseWriter, r *http.Request) {}

func (h handler) AddFood(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var food food.Food
    json.Unmarshal(body, &food)

    // Append to the Books table
    if result := h.DB.Create(&food); result.Error != nil {
        fmt.Println(result.Error)
    }

    // Send a 201 created response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}

func (h handler) UpdateFood(w http.ResponseWriter, r *http.Request) {}

func (h handler) DeleteFood(w http.ResponseWriter, r *http.Request) {
    // Read the dynamic id parameter
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // Find the food by Id

    var food food.Food

    if result := h.DB.First(&food, id); result.Error != nil {
        fmt.Println(result.Error)
    }

    // Delete that book
    h.DB.Delete(&food)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Deleted")
}
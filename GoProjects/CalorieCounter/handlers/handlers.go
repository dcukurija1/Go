package handlers

import (
    "gorm.io/gorm"
	"net/http"
)

type handler struct {
    DB *gorm.DB
}

func New(db *gorm.DB) handler {
    return handler{db}
}


func (h handler) GetAllFood(w http.ResponseWriter, r *http.Request) {}

func (h handler) GetFood(w http.ResponseWriter, r *http.Request) {}

func (h handler) AddFood(w http.ResponseWriter, r *http.Request) {}

func (h handler) UpdateFood(w http.ResponseWriter, r *http.Request) {}

func (h handler) DeleteFood(w http.ResponseWriter, r *http.Request) {}
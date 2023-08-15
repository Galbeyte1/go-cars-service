package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Galbeyte1/go-cars-service/internal/car"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to our Car Service
type Handler struct {
	Router *mux.Router
	// For comfortablility, adds method based routes
	// and path variables
	Service *car.Service
}

// Response - an object to store responses from our API
type Response struct {
	Message string
	Error string
}

// NewHandler - return a pointer to a Handler
func NewHandler(service *car.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	// Health Check Endpoint
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(w, Response{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})

	h.Router.HandleFunc("/api/car/{id}", h.GetCar).Methods("GET")
	h.Router.HandleFunc("/api/car/", h.PostCar).Methods("POST")
	h.Router.HandleFunc("/api/car/{id}", h.UpdateCar).Methods("PUT")
	h.Router.HandleFunc("/api/car/", h.GetAllCars).Methods("GET")

}

// GetCar - retrieve a car by ID
func (h *Handler) GetCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCar Endpoint Hit!")
}

// GetAllCars - retrieve all cars
func (h *Handler) GetAllCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllCars Endpoint Hit!")
}

// PostCar - add a new car
func (h *Handler) PostCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostCar Endpoint Hit!")
}

// UpdateCar - updates a car by ID
func (h *Handler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateCar Endpoint Hit!")
}



// Gracefully handle errors 
// sendOkResponse 
func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

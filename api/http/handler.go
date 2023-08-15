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
	vars := mux.Vars(r)
	id := vars["id"]

	car, err := h.Service.GetCar(id)
	if err != nil {
		sendErrorResponse(w, "Failed to retrieve car by car ID", err)
		return 
	}

	if err := sendOkResponse(w, car); err != nil {
		panic(err)
	}
}

// GetAllCars - retrieve all cars
func (h *Handler) GetAllCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllCars Endpoint Hit!")

	cars, err := h.Service.GetAllCars()
	if err != nil {
		sendErrorResponse(w, "Failed to retrieve all cars", err)
		return
	}

	if err := sendOkResponse(w, cars); err != nil {
		panic(err)
	}
}

// PostCar - add a new car
func (h *Handler) PostCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostCar Endpoint Hit!")
	var car car.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	car, err := h.Service.PostCar(car)

	if err != nil {
		sendErrorResponse(w, "Failed to post new car", err)
		return
	}
	if err := sendOkResponse(w, car); err != nil {
		panic(err)
	}
}

// UpdateCar - updates a car by ID
func (h *Handler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateCar Endpoint Hit!")
	var car car.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	car, err := h.Service.UpdateCar(id, car)
	if err != nil {
		sendErrorResponse(w, "Failed to update car", err)
		return
	}
	if err := sendOkResponse(w, car); err != nil {
		panic(err)
	}
}



// Gracefully handle errors 
// sendOkResponse 
func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}

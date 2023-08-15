package database

import (
	"github.com/Galbeyte1/go-cars-service/internal/car"
)

func NewDatabase() map[string]car.Car {
	db := make(map[string]car.Car)
	return db
}
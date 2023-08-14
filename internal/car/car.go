package car

import "errors"

// Service - the struct for our car service
type Service struct {
	DB map[string]Car
}

// Defines our car structure
type Car struct {
	ID string
	Make string
	Model string
	Package string
	Color string
	Year string
	Category string
	Mileage uint
	Price uint
}

// CarService - the interface for our car service
type CarService interface {
	GetCar(ID uint) (Car, error)
	GetAllCars() ([]Car, error)
	PostCar(car Car) (Car, error)
	UpdateCar(ID uint, newCar Car) (Car, error)
}

// NewService - returns a new car service
func NewService(db map[string]Car) *Service {
	return &Service{
		DB: db,
	}
}

// GetCar - retrieves car by their ID from database
func (s *Service) GetCar(ID string) (Car, error) {
	car, exist := s.DB[ID]
	if !exist {
		return Car{}, errors.New("car not found")
	}
	return car, nil
}

// GetAllCars - retrieves all cars from the database
func (s *Service) GetAllCars() ([]Car, error) {
	cars := make([]Car, 0, len(s.DB))
	for _, car := range s.DB {
		cars = append(cars, car)
	}
	return cars, nil
}

// PostCar - add a car to the database
func (s *Service) PostCar(car Car) (Car, error) {
	ID := car.ID
	s.DB[ID] = car
	return car, nil
}

// UpdateCar - update a car by ID in the database
func (s *Service) UpdateCar(ID string, newCar Car) (Car, error) {
	_, err := s.GetCar(ID)
	if err != nil {
		return Car{}, err
	}

	s.DB[ID] = newCar
	return newCar, err
}


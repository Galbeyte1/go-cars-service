package database

import "github.com/Galbeyte1/go-cars-service/internal/car"

func PopulateDB() map[string]car.Car {
	car1 := car.Car{
		ID:       "JHI290Xj",
		Make:     "Ford",
		Model:    "F10",
		Package:  "Base",
		Color:    "Silver",
		Year:     "2010",
		Category: "Truck",
		Mileage:  120123,
		Price:    1999900,
	}

	car2 := car.Car{
		ID:       "fW37la",
		Make:     "Toyota",
		Model:    "Camry",
		Package:  "SE",
		Color:    "White",
		Year:     "2019",
		Category: "Sedan",
		Mileage:  3999,
		Price:    2899000,
	}

	car3 := car.Car{
		ID:       "1i3xjRlIc",
		Make:     "Toyota",
		Model:    "Rav4",
		Package:  "XSE",
		Color:    "Red",
		Year:     "2018",
		Category: "SUV",
		Mileage:  24001,
		Price:    2275000,
	}

	car4 := car.Car{
		ID:       "dku43920s",
		Make:     "Ford",
		Model:    "Bronco",
		Package:  "Badlands",
		Color:    "Burnt Orange",
		Year:     "2022",
		Category: "SUV",
		Mileage:  1,
		Price:    4499000,
	}

	cars := make(map[string]car.Car)
	cars["JHK290Xj"] = car1
	cars["fW37la"] = car2
	cars["1i3xjRlIc"] = car3
	cars["dku43920s"] = car4

	return cars
}
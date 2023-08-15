//go:build ete
// +build ete

package car

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

const (
	BASE_URL = "http://localhost:8080"
)

// Unit Testing Car Service
func TestServiceGetCar(t *testing.T) {
	
}

func TestServiceGetAllCars(t *testing.T) {
	
}

func TestServicePostCar(t *testing.T) {
	
}

func TestServiceUpdateCar(t *testing.T) {
	
}

func TestGetCar(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/car/")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}

// func TestGetAllCars(t *testing.T) {
// 	client := resty.New()
// 	resp, err := client.R().Get(BASE_URL + "/api/car")
// 	if err != nil {
// 		t.Fail()
// 	}
// 	assert.Equal(t, 200, resp.StatusCode())
// }

// func TestPostCar(t *testing.T) {
// 	client := resty.New()
// 	resp, err := client.R().Get(BASE_URL + "/api/car")
// 	if err != nil {
// 		t.Fail()
// 	}
// 	assert.Equal(t, 200, resp.StatusCode())
// }

// func TestUpdateCar(t *testing.T) {
// 	client := resty.New()
// 	resp, err := client.R().Get(BASE_URL + "/api/car")
// 	if err != nil {
// 		t.Fail()
// 	}
// 	assert.Equal(t, 200, resp.StatusCode())
// }




package main

import (
	"cars-api/config"
	"cars-api/models"
	"cars-api/response"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func (env *Env) CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCar models.Car
	resp := response.New()
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if response.Check(err, w, resp, "", 100) {
		return
	}
	err = env.CarsManager.CarsDataStore.Validate(&newCar)
	if response.Check(err, w, resp, "", 101) {
		return
	}
	car := models.Car{
		ID:         "",
		Brand:      newCar.Brand,
		Model:      newCar.Model,
		Price:      newCar.Price,
		Status:     newCar.Status,
		Kilometres: newCar.Kilometres,
	}

	postedCar, err := env.CarsManager.PostCar(&car)
	if response.Check(err, w, resp, "failed to post car", 102) {
		return
	}

	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Car            *models.Car              `json:"car"`
	}{resp, postedCar})
	if response.Check(err, w, resp, "failed to marshal Json", 103) {
		return
	}
	w.Write(context)
	return
}

func (env *Env) GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := response.New()
	cars, err := env.CarsManager.CarsDataStore.GetAll()
	if response.Check(err, w, resp, "failed to get fetched records", 104) {
		return
	}
	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Cars           *[]models.Car            `json:"cars"`
	}{resp, cars})
	if response.Check(err, w, resp, "failed to marshal Json", 105) {
		return
	}
	w.Write(context)
	return

}

func (env *Env) UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := response.New()
	carID := chi.URLParam(r, "id")
	fetchedCar, err := env.CarsManager.CarsDataStore.Get(carID)
	if response.Check(err, w, resp, "", 106) {
		return
	}
	var car models.Car
	err = json.NewDecoder(r.Body).Decode(&car)
	if response.Check(err, w, resp, "failed to decode data", 107) {
		return
	}
	err = env.CarsManager.CarsDataStore.Validate(&car)
	if response.Check(err, w, resp, "", 108) {
		return
	}
	fetchedCar.Model = car.Model
	fetchedCar.Brand = car.Brand
	fetchedCar.Status = car.Status
	fetchedCar.Price = car.Price
	fetchedCar.Kilometres = car.Kilometres

	err = env.CarsManager.CarsDataStore.Update(fetchedCar)
	if response.Check(err, w, resp, "failed to update car", 109) {
		return
	}
	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Car            *models.Car              `json:"car"`
	}{resp, fetchedCar})
	if response.Check(err, w, resp, "failed to marshal Json", 110) {
		return
	}
	w.Write(context)
	return
}

func (env *Env) DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	carID := chi.URLParam(r, "id")
	resp := response.New()
	car, err := env.CarsManager.DeleteCar(carID)
	if response.Check(err, w, resp, "failed to delete fetched car", 111) {
		return
	}
	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Car            *models.Car              `json:"car"`
	}{resp, car})
	if response.Check(err, w, resp, "failed to marshal Json", 112) {
		return
	}
	w.Write(context)
	return
}

func (env *Env) GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	carID := chi.URLParam(r, "id")
	resp := response.New()
	car, err := env.CarsManager.CarsDataStore.Get(carID)
	if response.Check(err, w, resp, "failed to get fetched car", 113) {
		return
	}
	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Car            *models.Car              `json:"car"`
	}{resp, car})
	if response.Check(err, w, resp, "failed to Marshal Json", 113) {
		return
	}
	w.Write(context)
	return
}

func statusCheck(str string) bool {
	var statusMap = map[string]bool{
		config.AWAY_STATUS:                true,
		config.IN_STOCK_STATUS:            true,
		config.NO_LONGER_AVAILABLE_STATUS: true,
		config.SOLD_OUT_STATUS:            true,
	}
	if _, ok := statusMap[str]; !ok {
		return false
	} else {
		return true
	}
}

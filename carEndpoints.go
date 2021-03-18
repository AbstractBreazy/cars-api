package main

import (
	"cars-api/models"
	"cars-api/response"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"net/http"
)

func (env *Env) CreateCar(w http.ResponseWriter, r *http.Request) {
	type InitialData struct {
		Brand      string `json:"brand"`
		Model      string `json:"model"`
		Price      int64  `json:"price"`
		Status     string `json:"status"`
		Kilometres int64  `json:"kilometres"`
	}
	var newCar InitialData
	resp := response.New()
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if check(err, w, resp, "", 100) {
		return
	}
	if len(newCar.Brand) < 1 {
		if check(errors.New("incorrect Brand value"), w, resp, "", 101) {
			return
		}
	}
	if len(newCar.Model) < 1 {
		if check(errors.New("incorrect Model value"), w, resp, "", 102) {
			return
		}
	}
	if statusCheck(newCar.Status) == false {
		if check(errors.New("incorrect Status value"), w, resp, "", 103) {
			return
		}
	}
	if newCar.Kilometres < 0 {
		if check(errors.New("incorrect Kilometres value"), w, resp, "", 104) {
			return
		}
	}
	if newCar.Price < 0 {
		if check(errors.New("incorrect Price value"), w, resp, "", 105) {
			return
		}
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
	if check(err, w, resp, "failed to post car", 106) {
		return
	}

	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Car            *models.Car              `json:"car"`
	}{resp, postedCar})
	if check(err, w, resp, "failed to marshal Json", 107) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(context)
	return
}

func (env *Env) GetCars(w http.ResponseWriter, r *http.Request) {
	resp := response.New()
	cars, err := env.CarsManager.CarsDataStore.GetAll()
	if check(err, w, resp, "failed to get fetched records", 108) {
		return
	}
	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Cars           *[]models.Car            `json:"cars"`
	}{resp, cars})
	if check(err, w, resp, "failed to marshal Json", 109) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(context)
	return

}

func (env *Env) UpdateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	resp := response.New()
	err := json.NewDecoder(r.Body).Decode(&car)
	if check(err, w, resp, "failed to decode data", 110) {
		return
	}
	oldCar, err := env.CarsManager.CarsDataStore.Get(car.ID)
	if check(err, w, resp, "failed to get fetched car", 111) {
		return
	}
	if len(car.Model) < 1 {
		if check(errors.New("incorrect Model value"), w, resp, "", 112) {
			return
		}
	}
	if len(car.Brand) < 1 {
		if check(errors.New("incorrect Brand value"), w, resp, "", 113) {
			return
		}
	}
	if statusCheck(car.Status) == false {
		if check(errors.New("incorrect Status value"), w, resp, "", 114) {
			return
		}
	}
	if car.Price < 0 {
		if check(errors.New("incorrect Price value"), w, resp, "", 115) {
			return
		}
	}
	if car.Kilometres < 0 {
		if check(errors.New("incorrect Kilometres value"), w, resp, "", 116) {
			return
		}
	}
	oldCar.Model = car.Model
	oldCar.Brand = car.Brand
	oldCar.Kilometres = car.Kilometres
	oldCar.Price = car.Price
	oldCar.Status = car.Status
	err = env.CarsManager.CarsDataStore.Update(oldCar)
	if check(err, w, resp, "failed to update car", 117) {
		return
	}
	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Car            *models.Car              `json:"car"`
	}{resp, oldCar})
	if check(err, w, resp, "failed to marshal Json", 118) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(context)
	return
}

func (env *Env) DeleteCar(w http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "id")
	resp := response.New()
	car, err := env.CarsManager.DeleteCar(carID)
	if check(err, w, resp, "failed to delete fetched car", 119) {
		return
	}
	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Car            *models.Car              `json:"car"`
	}{resp, car})
	if check(err, w, resp, "failed to marshal Json", 120) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(context)
	return
}

func (env *Env) GetCar(w http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "id")
	resp := response.New()
	car, err := env.CarsManager.CarsDataStore.Get(carID)
	if check(err, w, resp, "failed to get fetched car", 121) {
		return
	}
	context, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Car            *models.Car              `json:"car"`
	}{resp, car})
	if check(err, w, resp, "failed to Marshal Json", 122) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(context)
	return
}

func check(err error, w http.ResponseWriter, resp *response.StatusResponse, message string, code int) bool {
	if err != nil {
		resp.Status = false
		if len(message) == 0 {
			resp.ResponseMessage = err.Error()
		} else {
			resp.ResponseMessage = message + ": " + err.Error()
		}
		resp.ResponseInternalCode = code
		w.Write(resp.GetJSON())
		return true
	}
	return false
}

func statusCheck(str string) bool {
	expectedStatusList := []string{
		"Away",
		"In Stock",
		"Sold Out",
		"No longer available",
	}
	for _, b := range expectedStatusList {
		if b == str {
			return true
		}
	}
	return false
}

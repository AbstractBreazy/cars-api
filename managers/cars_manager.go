package managers

import (
	"cars-api/models"
	uuid "github.com/satori/go.uuid"
)

type CarsManager struct {
	CarsDataStore models.CarsDataStore
}

func InitCarsManager(connection *models.Connection) CarsManager {
	return CarsManager{
		CarsDataStore: models.CarsDataStore{connection}}
}

func (m *CarsManager) PostCar(car *models.Car) (*models.Car, error) {
	carID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	car.ID = carID.String()
	err = m.CarsDataStore.Insert(car)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (m *CarsManager) DeleteCar(carId string) (*models.Car, error) {
	car, err := m.CarsDataStore.Get(carId)
	if err != nil {
		return nil, err
	}
	err = m.CarsDataStore.Remove(car.ID)
	if err != nil {
		return nil, err
	}
	return car, nil
}

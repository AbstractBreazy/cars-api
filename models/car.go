package models

import (
	"cars-api/config"
	"errors"
)

type Car struct {
	ID         string `json:"id" gorm:"primary_key"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Price      int64  `json:"price"`
	Status     string `json:"status"` // 1 - Away, 2 - In Stock, 3 - Sold Out, 4 - No longer available
	Kilometres int64  `json:"kilometres"`
}

type CarsDataStore struct {
	*Connection
}

func (b *CarsDataStore) Insert(obj *Car) error {
	err := b.Connection.DB.Create(&obj).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *CarsDataStore) Update(obj *Car) error {
	err := b.Connection.DB.Model(&obj).Save(&obj).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *CarsDataStore) Get(id string) (*Car, error) {
	car := &Car{}
	err := b.Connection.DB.Where("id = ?", id).First(&car).Error
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (b *CarsDataStore) Remove(id string) error {
	car := &Car{}
	err := b.Connection.DB.Where("id = ?", id).Delete(&car).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *CarsDataStore) GetAll() (*[]Car, error) {
	cars := &[]Car{}
	err := b.Connection.DB.Find(&cars).Error
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (b *CarsDataStore) Validate(obj *Car) (err error) {
	if len(obj.Model) < 1 {
		return errors.New("incorrect model value")
	}
	if len(obj.Brand) < 1 {
		return errors.New("incorrect brand value")
	}
	if StatusCheck(obj.Status) == false {
		return errors.New("incorrect status value")
	}
	if obj.Price < 0 {
		return errors.New("incorrect price value")
	}
	if obj.Kilometres < 0 {
		return errors.New("incorrect kilometres value")
	}
	return nil
}

func StatusCheck(str string) bool {
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

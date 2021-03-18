package models

/*
Тестовое задание:

Разработать CRUD (REST API) для модели автомобиля, который имеет следующие поля:



1. Уникальный идентификатор (любой тип, общение с БД не является критерием чего-либо, можно сделать и in-memory хранилище на время жизни сервиса)

2. Бренд автомобиля (текст)

3. Модель автомобиля (текст)

4. Цена автомобиля (целое, не может быть меньше 0)

5. Статус автомобиля (В пути, На складе, Продан, Снят с продажи)

6. Пробег (целое)

*/
type Car struct {
	ID         string `json:"id"` // unique / primary key gorm
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Price      int64  `json:"price"`
	Status     string `json:"status"` // 1 - В пути, 2 - На Складе, 3 - Продано, 4 - Нет в наличии
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

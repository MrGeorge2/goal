package seeder

import (
	"math/rand"
	"time"

	"github.com/MrGeorge2/goal/list"
)

type Car struct {
	ID    uint64
	Brand string
	IsNew bool
}

func CreateOrderedCarList() list.List[Car] {
	cars := list.List[Car]{}

	cars = append(cars, Car{
		ID:    1,
		Brand: "Volvo",
		IsNew: true,
	})

	cars = append(cars, Car{
		ID:    2,
		Brand: "Mercedes",
		IsNew: true,
	})

	cars = append(cars, Car{
		ID:    3,
		Brand: "Skoda",
		IsNew: false,
	})

	cars = append(cars, Car{
		ID:    4,
		Brand: "Skoda",
		IsNew: true,
	})

	cars = append(cars, Car{
		ID:    5,
		Brand: "BMW",
		IsNew: true,
	})

	return cars
}

func CreateShuffledCarList() list.List[Car] {
	carList := CreateOrderedCarList()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(carList), func(i, j int) { carList[i], carList[j] = carList[j], carList[i] })
	return carList
}

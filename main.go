package main

import (
	"fmt"

	"github.com/MrGeorge2/goal/tests/seeder"
)

func main() {
	carList := seeder.CreateShuffledCarList()
	//originalLen := len(carList)

	lada := seeder.Car{
		Brand: "Lada",
		ID:    100,
		IsNew: false,
	}

	carList.Add(lada)
	fmt.Println(carList)
}

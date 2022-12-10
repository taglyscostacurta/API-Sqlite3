package main

import (
	"database/sql" // database driver

	"github.com/labstack/echo/v4"   // web framework
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

type Car struct {
	Name  string  `json:"name"` // json tag
	Price float64 `json:"price"`
}

var cars []Car // slice of cars

func generateCars() { // generate some cars
	cars = append(cars, Car{"Ferrari", 100000})
	cars = append(cars, Car{"Lamborghini", 200000})
	cars = append(cars, Car{"Porsche", 300000})
}

func main() {

	generateCars()                   // generate some cars
	e := echo.New()                  // create new echo instance
	e.GET("/cars", getCars)          // register handler
	e.POST("/cars", createCar)       // register handler
	e.Logger.Fatal(e.Start(":8000")) // start server

}

func getCars(c echo.Context) error { // handler
	return c.JSON(200, cars) // return json
}

func createCar(c echo.Context) error {
	car := new(Car) // create new car
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(201, car)
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db") // open database
	if err != nil {
		return err
	}
	defer db.Close() // close database

	stmt, err := db.Prepare("INSERT INTO cars(name, price) values($1, $2)") // prepare statement
	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name, car.Price) // execute statement
	if err != nil {
		return err
	}
	return nil
}

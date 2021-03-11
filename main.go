package main

import (
	"encoding/json"

	"log"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"

	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Driver struct {
	gorm.Model

	Name string

	License string

	Data string

	Cars []Car
}

type Car struct {
	gorm.Model

	Year int

	Make string

	ModelName string

	DriverID int
}

var db *gorm.DB

var err error

var (
	drivers = []Driver{

		{Name: "Lady Bug", License: "DBA123", Data: "Hello from the other side!"},

		{Name: "Dude Bug", License: "DEV789", Data: "Hello Wolrd"},

		{Name: "Pepa Pig", License: "OPS333", Data: "Wolrd!!! Hello!!!"},
	}

	cars = []Car{

		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},

		{Year: 2018, Make: "Honda", ModelName: "Accord", DriverID: 1},

		{Year: 2015, Make: "Nissan", ModelName: "Sentra", DriverID: 2},

		{Year: 2019, Make: "Volks", ModelName: "New Beatle", DriverID: 3},
	}
)

func main() {

	router := mux.NewRouter()

	db, err = gorm.Open("postgres", "host=172.18.0.2 port=5432 user=postgres dbname=go-auth sslmode=disable password=postgres")

	if err != nil {

		panic("failed to connect database")

	}

	defer db.Close()

	db.AutoMigrate(&Driver{})

	db.AutoMigrate(&Car{})

	for index := range cars {

		db.Create(&cars[index])

	}

	for index := range drivers {

		db.Create(&drivers[index])

	}

	router.HandleFunc("/data", GetData).Methods("GET")

	router.HandleFunc("/cars", GetCars).Methods("GET")

	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")

	router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")

	router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

func GetData(w http.ResponseWriter, r *http.Request) {

	var data []Driver

	db.Find(&data)

	json.NewEncoder(w).Encode(&data)

}

func GetCars(w http.ResponseWriter, r *http.Request) {

	var cars []Car

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}

func GetCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var car Car

	db.First(&car, params["id"])

	json.NewEncoder(w).Encode(&car)

}

func GetDriver(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var driver Driver

	var cars []Car

	db.First(&driver, params["id"])

	db.Model(&driver).Related(&cars)

	driver.Cars = cars

	json.NewEncoder(w).Encode(&driver)

}

func DeleteCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var car Car

	db.First(&car, params["id"])

	db.Delete(&car)

	var cars []Car

	db.Find(&cars)

	json.NewEncoder(w).Encode(&cars)

}

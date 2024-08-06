package main

import "fmt"

type Vehicle struct {
	//model string
	model  ModelStruct
	volume float32
}

func (vehicle Vehicle) getVolume() float32 {
	return vehicle.volume
}

func (vehicle *Vehicle) setVolume(volume float32) {
	vehicle.volume = volume
}

type Car struct {
	//model   string
	model ModelStruct
	//Vehicle
	vehicle Vehicle
}

func (car Car) getVehicle() Vehicle {
	return car.vehicle
}

func (car *Car) setVehicle(vehicle Vehicle) {
	car.vehicle = vehicle
}

type ModelInterface interface {
	getModel() string
	setModel(model string)
}

type ModelStruct struct {
	model string
}

//func getModel(modelInterface ModelInterface) string {
//	return modelInterface.getModel()
//}

//func (modelInterface *ModelInterface) modelChangeField2nd(field1st string, inputData string) {
//	// field2nd = ...
//}

func (modelStruct ModelStruct) getModel() string {
	return modelStruct.model
}

func (modelStruct *ModelStruct) setModel(model string) {
	modelStruct.model = model
}

func CAR_INFO(car Car) {
	fmt.Println("car model:", car.model)
}

func FULL_INFO(car Car) {
	fmt.Println("FULL INFO CAR:")
	CAR_INFO(car)
	VEHICLE_INFO(car.getVehicle())
}

func VEHICLE_INFO(vehicle Vehicle) {
	fmt.Println("volume:", vehicle.getVolume())
	fmt.Println("vehicle model:", vehicle.model.getModel())
}

func main() {
	//car := Car{Vehicle: Vehicle{model: "diesel", volume: 4.8}, model: "bmw"}
	//car := Car{vehicle: Vehicle{model: "diesel", volume: 4.8}, model: "bmw"}
	//fmt.Println("car:", car)

	var car = Car{model: ModelStruct{model: "bmw"},
		vehicle: Vehicle{model: ModelStruct{"diesel"}, volume: 4.8}}
	fmt.Println("car:", car)

	type CarChange interface {
		getVehicle() Vehicle
		setVehicle(vehicle Vehicle)
	}

	//var carChange CarChange = &car
	//VEHICLE_INFO(carChange.getVehicle())
	//
	//var car1 = car
	//CAR_INFO(car1)

	FULL_INFO(car)
}

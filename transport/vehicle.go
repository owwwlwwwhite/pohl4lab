package transport

import "fmt"

// Vehicle - базовый тип транспорта
type Vehicle struct {
	brand       string
	year        int
	vehicleType string
}

// NewVehicle создает новый экземпляр Vehicle
func NewVehicle(brand string, year int, vehicleType string) *Vehicle {
	return &Vehicle{
		brand:       brand,
		year:        year,
		vehicleType: vehicleType,
	}
}

func (v *Vehicle) StartEngine() {
	fmt.Println(v.vehicleType + " " + v.brand + " заводится.")
}

func (v *Vehicle) Move() {
	fmt.Println(v.vehicleType + " " + v.brand + " движется.")
}

func (v *Vehicle) Stop() {
	fmt.Println(v.vehicleType + " " + v.brand + " останавливается.")
}

func (v *Vehicle) GetBrand() string {
	return v.brand
}

func (v *Vehicle) GetYear() int {
	return v.year
}

func (v *Vehicle) GetVehicleType() string {
	return v.vehicleType
}

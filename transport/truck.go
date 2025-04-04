package transport

import (
	"fmt"
	"math/rand"
	"time"
)

// Truck представляет грузовик, использующий Vehicle
type Truck struct {
	*Vehicle       // Встраивание Vehicle (композиция)
	cargoCapacity int
}

// NewTruck создает новый экземпляр Truck
func NewTruck(brand string, year int, vehicleType string, cargoCapacity int) *Truck {
	rand.Seed(time.Now().UnixNano())
	
	return &Truck{
		Vehicle:       NewVehicle(brand, year, vehicleType), // Используем конструктор Vehicle
		cargoCapacity: cargoCapacity,
	}
}

// LoadCargo имитирует загрузку груза
func (t *Truck) LoadCargo() {
	success := rand.Intn(2)
	if success == 1 {
		fmt.Printf("%s %s успешно загрузил груз!\n", t.GetVehicleType(), t.GetBrand())
	} else {
		fmt.Printf("%s %s не смог загрузить груз из-за перевеса.\n", t.GetVehicleType(), t.GetBrand())
	}
}

// GetCargoCapacity возвращает грузоподъемность
func (t *Truck) GetCargoCapacity() int {
	return t.cargoCapacity
}

// StartTruckEngine специфичный метод для Truck
func (t *Truck) StartTruckEngine() {
	fmt.Printf("Грузовик %s готов к работе!\n", t.GetBrand())
	t.StartEngine() // Используем метод Vehicle
}

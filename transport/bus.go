package transport

import (
	"fmt"
	"math/rand"
	"time"
)

// Bus представляет автобус, использующий Truck
type Bus struct {
	*Truck              // Встраивание Truck
	passengerCapacity int
}

// NewBus создает новый экземпляр Bus
func NewBus(brand string, year int, vehicleType string, cargoCapacity int, passengerCapacity int) *Bus {
	rand.Seed(time.Now().UnixNano())
	
	return &Bus{
		Truck:             NewTruck(brand, year, vehicleType, cargoCapacity), // Используем конструктор Truck
		passengerCapacity: passengerCapacity,
	}
}

// AnnounceStop объявляет остановку
func (b *Bus) AnnounceStop() {
	announcements := []string{
		"Осторожно, двери закрываются!",
		"Следующая остановка - Центральная площадь",
		"Пожалуйста, оплатите проезд",
	}
	fmt.Printf("%s %s объявляет: %s\n", b.GetVehicleType(), b.GetBrand(), announcements[rand.Intn(len(announcements))])
}

// TransportPassengers имитирует перевозку пассажиров
func (b *Bus) TransportPassengers() {
	fmt.Printf("%s %s перевозит %d пассажиров.\n", b.GetVehicleType(), b.GetBrand(), b.passengerCapacity)
}

// GetPassengerCapacity возвращает вместимость пассажиров
func (b *Bus) GetPassengerCapacity() int {
	return b.passengerCapacity
}

// StartBusEngine специфичный метод для Bus
func (b *Bus) StartBusEngine() {
	fmt.Printf("Автобус %s готов к маршруту!\n", b.GetBrand())
	b.StartEngine() // Используем метод Vehicle через Truck
}

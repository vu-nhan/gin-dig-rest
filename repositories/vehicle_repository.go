package repositories

type vehicleRepository struct {

}

type VehicleRepository interface {

}

func NewVehicleRepository() VehicleRepository {
	return vehicleRepository{}
}

package services

type vehicleService struct {

}

type VehicleService interface {

}

func NewVehicleService() VehicleService {
	return &vehicleService{}
}

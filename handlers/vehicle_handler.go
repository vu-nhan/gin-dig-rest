package handlers

import (
	"github.com/vu-nhan/gin-dig-rest/services"
)

type VehicleHandler struct {
	vehicleService services.VehicleService
}

func NewVehicleHandler(injectedVehicleService services.VehicleService) *VehicleHandler {
	return &VehicleHandler{
		vehicleService: injectedVehicleService,
	}
}

func (h * VehicleHandler) GetAllVehicle() {

}
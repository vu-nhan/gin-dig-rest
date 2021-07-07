package provider

import (
	"github.com/vu-nhan/gin-dig-rest/handlers"
	"github.com/vu-nhan/gin-dig-rest/services"
	"go.uber.org/dig"
)

type IocProvider struct {
	DigContainer *dig.Container
}

func (c *IocProvider) InitIocProvider() {
	c.DigContainer = dig.New()

	_ = c.DigContainer.Provide(services.NewVehicleService)
	_ = c.DigContainer.Provide(handlers.NewVehicleHandler)
}

func (c *IocProvider) Invoke(injectedInterface interface{}) error {
	return c.DigContainer.Invoke(injectedInterface)
}

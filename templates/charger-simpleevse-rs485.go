package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "simpleevse",
		Name:   "EVSE DIN (Modbus RTU)",
		Sample: `device: /dev/ttyUSB0 # serial RS485 interface`,
	}

	registry.Add(template)
}

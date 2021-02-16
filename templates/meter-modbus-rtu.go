package templates

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class: "meter",
		Type:  "modbus",
		Name:  "Eastron SDM Modbus Meter (RTU)",
		Sample: `model: sdm
device: /dev/ttyUSB0 # serial port
id: 2`,
	}

	registry.Add(template)
}

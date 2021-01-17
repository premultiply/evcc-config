package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Eastron SDM Modbus Meter ('RTU-over-TCP')",
		Sample: `model: sdm
uri: 192.0.2.2:502
rtu: true # rs485 device connected using simple serial/tcp gateway
id: 2
power: Power # default values, optionally override
energy: Sum # default values, optionally override`,
	}

	registry.Add(template)
}

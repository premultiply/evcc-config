package templates

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class: "meter",
		Type:  "modbus",
		Name:  "Generic SunSpec Battery inverter (Battery Meter)",
		Sample: `uri: 192.168.178.100:502
id: 1
soc: ChargeState`,
	}

	registry.Add(template)
}

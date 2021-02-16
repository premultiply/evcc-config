package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "SMA SunnyBoy / TriPower / other PV-inverter",
		Sample: `uri: 192.168.178.100:502
id: 126`,
	}

	registry.Add(template)
}

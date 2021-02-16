package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Kostal Smart Energy Meter (Grid Meter)",
		Sample: `uri: 192.168.178.100:502
id: 71`,
	}

	registry.Add(template)
}

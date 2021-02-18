package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "simpleevse",
		Name:   "EVSE DIN (Modbus/TCP)",
		Sample: `uri: 192.0.2.2:502 # Modbus/TCP converter adress`,
	}

	registry.Add(template)
}

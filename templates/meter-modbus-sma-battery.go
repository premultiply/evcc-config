package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "SMA Sunny Island / Sunny Boy Storage",
		Sample: `uri: 192.168.178.100:502
id: 126
soc: ChargeState`,
	}

	registry.Add(template)
}

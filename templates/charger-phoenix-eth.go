package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "phoenix-eth",
		Name:   "Phoenix ETH Controller (Modbus/TCP)",
		Sample: `uri: 192.0.2.2:502
id: 255 # Modbus id, may be 180 or 255 (default)`,
	}

	registry.Add(template)
}

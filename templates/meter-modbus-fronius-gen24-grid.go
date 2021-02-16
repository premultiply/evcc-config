package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Fronius Symo GEN24 Plus (Grid Meter)",
		Sample: `model: 213 # sunspec meter
uri: 192.0.2.2:502
id: 200`,
	}

	registry.Add(template)
}

package templates

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class: "meter",
		Type:  "modbus",
		Name:  "Generic SunSpec 3-phase meter (Grid Meter)",
		Sample: `model: 203 # Wye Connect Meter Model
uri: 192.0.2.2:502
id: 1`,
	}

	registry.Add(template)
}

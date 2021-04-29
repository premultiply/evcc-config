package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Multiple PV inverters combined (PV Meter)",
		Sample: `power:
  plugin: calc
  add:
  - plugin: modbus
    model: sunspec
    uri: 192.0.2.2:502
    id: 1
  - plugin: modbus
    model: sunspec
    uri: 192.0.2.3:502
    id: 1`,
	}

	registry.Add(template)
}

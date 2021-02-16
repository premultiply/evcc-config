package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "Multiple PV inverters combined (PV Meter)",
		Sample: `power:
  type: calc
  add:
  - type: modbus
    model: sunspec
    uri: 192.168.178.101:502
    id: 1
  - type: modbus
    model: sunspec
    uri: 192.168.178.102:502
    id: 1`,
	}

	registry.Add(template)
}

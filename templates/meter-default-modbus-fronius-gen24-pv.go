package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Fronius Symo GEN24 Plus (PV Meter)",
		Sample: `power:
  plugin: calc
  add:
  - plugin: modbus
    model: sunspec
    uri: 192.0.2.2:502
    id: 1
    value: 160:1:DCW # mpp 1 pv
  - plugin: modbus
    model: sunspec
    uri: 192.0.2.2:502
    id: 1
    value: 160:2:DCW # mpp 2 pv`,
	}

	registry.Add(template)
}

package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "SolarEdge Hybrid Inverter (PV Meter)",
		Sample: `power:
  plugin: calc
  add:
  - plugin: modbus
    model: sunspec
    uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
    id: 1
    value: 103:DCW
  - plugin: modbus
    uri: 192.0.2.2:502 # Port 502 (SetApp) or 1502 (LCD)
    id: 1
    register:
      address: 62836 # Battery 1 Instantaneous Power
      type: holding
      decode: float32s`,
	}

	registry.Add(template)
}

package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "default",
		Name:   "SMA SHM or EM via inverter (Grid Meter)",
		Sample: `power:
  type: calc
  add:
  - type: modbus
    uri: 192.0.2.2:502
    id: 3
    register: # manual non-sunspec register configuration
      address: 30865 # SMA Modbus Profile: Power grid reference
      type: holding
      decode: int32
  - type: modbus
    uri: 192.0.2.2:502
    id: 3
    register: # manual non-sunspec register configuration
      address: 30867 # SMA Modbus Profile: Power grid feed-in
      type: holding
      decode: int32
    scale: -1`,
	}

	registry.Add(template)
}

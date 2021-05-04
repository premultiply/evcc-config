package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "E3DC (Battery Meter)",
		Sample: `power:
  source: modbus
  uri: e3dc.fritz.box:502
  id: 1 # ModBus slave id
  register: # manual register configuration
    address: 40069
    type: holding
    decode: int32s
  scale: -1 # reverse direction
soc:
  source: modbus
  uri: e3dc.fritz.box:502
  id: 1 # ModBus slave id
  register: # manual register configuration
    address: 40082
    type: holding
    decode: uint16`,
	}

	registry.Add(template)
}

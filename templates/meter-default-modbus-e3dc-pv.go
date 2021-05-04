package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "E3DC (PV Meter)",
		Sample: `power:
  source: modbus
  uri: e3dc.fritz.box:502
  id: 1 # ModBus slave id
  register: # manual register configuration
    address: 40067 # (40068 - 1) "Photovoltaikleistung in Watt"
    type: holding
    decode: int32s`,
	}

	registry.Add(template)
}

package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "default",
		Name:   "i-CHARGE CION (Modbus RTU-over-TCP)",
		Sample: `status:
  type: modbus
  uri: 192.0.2.2:502
  rtu: true # Modbus over TCP
  id: 1
  register: # manual register configuration
      address: 139 # CP-Status
      type: holding
      decode: uint16
enabled:
  type: modbus
  uri: 192.0.2.2:502
  rtu: true # Modbus over TCP
  id: 1 
  register: # manual register configuration
    address: 100 # Zustand
    type: holding
    decode: uint16
enable:
  type: modbus
  uri: 192.0.2.2:502
  rtu: true # Modbus over TCP
  id: 1
  register: # manual register configuration
    address: 100 # ein / aus
    type: writesingle
    decode: uint16
maxcurrent:
  type: modbus
  uri: 192.0.2.2:502
  rtu: true # Modbus over TCP
  id: 1
  register: # manual register configuration
    address: 127 # Strom max
    type: writesingle
    decode: uint16`,
	}

	registry.Add(template)
}

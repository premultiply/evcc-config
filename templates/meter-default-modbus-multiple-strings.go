package templates

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class: "meter",
		Type:  "default",
		Name:  "Multiple DC MPP strings combined (PV Meter)",
		Sample: `power:
  type: calc # use the calc plugin
  add: # The add function sums up both string values
  - type: modbus
    model: sunspec
    value: 160:1:DCW # MPP string 1
    uri: 192.168.178.101:502
    id: 1
  - type: modbus
    model: sunspec
    value: 160:2:DCW # MPP string 2
    uri: 192.168.178.101:502
    id: 1`,
	}

	registry.Add(template)
}

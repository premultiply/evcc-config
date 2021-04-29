package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "custom",
		Name:   "Generic",
		Sample: `status: # charger status A..F
  plugin: ...
  # ...
enabled: # charger enabled state (true/false or 0/1)
  plugin: ...
  # ...
enable: # set charger enabled state
  plugin: ...
  # ...
maxcurrent: # set charger max current
  plugin: ...
  # ...`,
	}

	registry.Add(template)
}

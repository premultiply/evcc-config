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
  source: ...
  # ...
enabled: # charger enabled state (true/false or 0/1)
  source: ...
  # ...
enable: # set charger enabled state
  source: ...
  # ...
maxcurrent: # set charger max current
  source: ...
  # ...`,
	}

	registry.Add(template)
}

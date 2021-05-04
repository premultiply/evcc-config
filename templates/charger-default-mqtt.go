package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "custom",
		Name:   "Generic (MQTT)",
		Sample: `status: # charger status A..F
  source: mqtt
  topic: some/topic1
enabled: # charger enabled state (true/false or 0/1)
  source: mqtt
  topic: some/topic2
enable: # set charger enabled state
  source: script
  cmd: /bin/sh -c "echo ${enable}"
maxcurrent: # set charger max current
  source: script
  cmd: /bin/sh -c "echo ${maxcurrent}"`,
	}

	registry.Add(template)
}

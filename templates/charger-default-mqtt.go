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
  plugin: mqtt
  topic: some/topic1
enabled: # charger enabled state (true/false or 0/1)
  plugin: mqtt
  topic: some/topic2
enable: # set charger enabled state
  plugin: script
  cmd: /bin/sh -c "echo ${enable}"
maxcurrent: # set charger max current
  plugin: script
  cmd: /bin/sh -c "echo ${maxcurrent}"`,
	}

	registry.Add(template)
}

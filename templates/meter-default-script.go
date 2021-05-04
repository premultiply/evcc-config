package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Generic (Script)",
		Sample: `power:
  source: script # use script plugin
  cmd: /bin/sh -c "echo 0" # actual command
  timeout: 3s # kill script after 3 seconds`,
	}

	registry.Add(template)
}

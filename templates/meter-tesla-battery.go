package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "tesla",
		Name:   "Tesla Powerwall (Battery)",
		Sample: `uri: http://192.168.178.100/
usage: battery`,
	}

	registry.Add(template)
}

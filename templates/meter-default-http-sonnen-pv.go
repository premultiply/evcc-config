package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Sonnenbatterie Eco/10 (PV Meter/ HTTP)",
		Sample: `power:
  source: http
  uri: http://192.0.2.2:8080/api/v1/status
  jq: .Production_W`,
	}

	registry.Add(template)
}

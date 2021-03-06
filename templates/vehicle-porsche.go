package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "porsche",
		Name:   "Porsche",
		Sample: `title: Taycan # display name for UI
capacity: 83 # kWh
user: # user
password: # password
vin: WP...
cache: 5m # cache API response`,
	}

	registry.Add(template)
}

package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "audi",
		Name:   "Audi",
		Sample: `title: Q55 TFSIe # display name for UI
capacity: 10 # kWh
user: # user
password: # password
vin: WAUZZZ...
cache: 5m # cache API response`,
	}

	registry.Add(template)
}

package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "tesla",
		Name:   "Tesla",
		Sample: `title: Model S # display name for UI
capacity: 90 # kWh
clientid: # client id
clientsecret: # client secret
email: # email
password: # password
vin: WTSLA...
cache: 5m # cache API response`,
	}

	registry.Add(template)
}

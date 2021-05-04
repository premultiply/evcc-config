package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Fronius Solar API V1 (Grid Meter/ HTTP)",
		Sample: `power:
  source: http
  uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
  jq: if .Body.Data.Site.P_Grid == null then 0 else .Body.Data.Site.P_Grid end`,
	}

	registry.Add(template)
}

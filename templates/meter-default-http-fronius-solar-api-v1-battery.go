package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "Fronius Solar API V1 (Battery Meter/ HTTP)",
		Sample: `power:
  source: http
  uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
  jq: if .Body.Data.Site.P_Akku == null then 0 else .Body.Data.Site.P_Akku end
soc:
  source: http
  uri: http://192.0.2.2/solar_api/v1/GetPowerFlowRealtimeData.fcgi
  jq: .Body.Data.Inverters."1".SOC`,
	}

	registry.Add(template)
}

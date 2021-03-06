package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "default",
		Name:   "OpenWB (remote-controlled using MQTT)",
		Sample: `status:
  # with openWB, charging status (A..F) this is split between "plugged" and "charging"
  # the openwb plugin combines both into status (charging=C, plugged=B, otherwise=A)
  type: openwb # use openwb plugin
  plugged:
    type: mqtt
    topic: openWB/lp/1/boolPlugStat
  charging:
    type: mqtt
    topic: openWB/lp/1/boolChargeStat
enabled:
  type: mqtt
  topic: openWB/lp/1/ChargePointEnabled
  timeout: 30s
enable:
  type: mqtt
  topic: openWB/set/lp1/ChargePointEnabled
  payload: ${enable:%d} # write payload definition
maxcurrent:
  type: mqtt
  topic: openWB/set/lp1/DirectChargeAmps
  payload: ${maxCurrent:%d} # write payload definition`,
	}

	registry.Add(template)
}

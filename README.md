# Configuration examples for EVCC

[![Build Status](https://travis-ci.org/andig/evcc-config.svg?branch=master)](https://travis-ci.org/andig/evcc-config)

Configuration examples for the [EVCC EV Charge Controller](https://github.com/andig/evcc).

## Meters

- [Eastron SDMxxx (Modbus over TCP)](#meter-modbus-sdm)
- [E3DC (Battery Meter)](#meter-e3dc-battery-meter)
- [E3DC (PV Meter)](#meter-e3dc-pv-meter)
- [Kostal Inverter (Grid Meter)](#meter-kostal-inverter-grid-meter)
- [Kostal Inverter (PV Meter)](#meter-kostal-inverter-pv-meter)
- [Kostal Smart Energy Meter (Grid Meter)](#meter-kostal-smart-energy-meter-grid-meter)
- [SMA (or other SunSpec compatible) Inverter (PV Meter)](#meter-sma-sunspec-pv-meter)
- [Multiple SunSpec compatible Inverters combined (PV Meter)](#meter-multiple-inverters-combined-pv-meter)
- [SMA Home Manager 2.0 / SMA Energy Meter](#meter-sma-home-manager-2-0--sma-energy-meter)
- [SolarLog (Grid Meter)](#meter-solarlog-grid-meter)
- [SolarLog (PV Meter)](#meter-solarlog-pv-meter)
- [vzlogger (HTTP)](#meter-vzlogger-http)
- [vzlogger (Push Server/ Websocket)](#meter-vzlogger-push-server-websocket)
- [vzlogger (split import/export channels)](#meter-vzlogger-split-import-export-channels)
- [Generic (MQTT)](#meter-generic-mqtt)
- [Generic (Script)](#meter-generic-script)

## Chargers

- [go-eCharger (Cloud)](#charger-go-echarger-cloud)
- [go-eCharger (Lokal)](#charger-go-echarger-lokal)
- [KEBA Connect](#charger-keba-connect)
- [Mobile Charger Connect](#charger-mobile-charger-connect)
- [NRGKick Bluetooth (direct)](#charger-nrgkick-bluetooth)
- [NRGKick via Connect device](#charger-nrgkick-connect)
- [OpenWB (remote-controlled using MQTT)](#charger-openwb-remote-controlled-using-mqtt)
- [Phoenix EM-CP Controller (Modbus-TCP)](#charger-phoenix-em-cp-controller-modbus-tcp)
- [Phoenix EV-CC Controller (Modbus)](#charger-phoenix-ev-cc-controller-modbus)
- [Simple EVSE (Modbus over TCP)](#charger-simple-evse-tcp)
- [Simple EVSE (Modbus on USB)](#charger-simple-evse-usb)
- [EVSE Wifi](#charger-evse-wifi)
- [Wallbe (Eco, Pro)](#charger-wallbe-eco-pro)
- [Wallbe (pre 2019 EV-CC-AC1 controller)](#charger-wallbe-pre-2019-ev-cc-ac1-controller)
- [Generic (MQTT)](#charger-generic-mqtt)
- [Generic (Script)](#charger-generic-script)

## Vehicles

- [Audi](#vehicle-audi)
- [BMW (i3)](#vehicle-bmw-i3)
- [Nissan (Leaf)](#vehicle-nissan-leaf)
- [Porsche](#vehicle-porsche)
- [Renault (Zoe)](#vehicle-renault-zoe)
- [Tesla](#vehicle-tesla)
- [Generic (Script)](#vehicle-generic-script)

## Details

### Meters

<a id="meter-modbus-sdm"></a>
#### Eastron SDMxxx (Modbus over TCP)

```yaml
- type: modbus
  model: sdm
  uri: 192.168.50.2:5002
  rtu: true # Modbus-RTU (not native Modbus-TCP)
  id: 2 # Modbus Device ID
  power: Power # default values, optionally override
  energy: Sum # default values, optionally override
```

<a id="meter-e3dc-battery-meter"></a>
#### E3DC (Battery Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: e3dc.fritz.box:502
    id: 1 # Modbus Device ID
    register: # manual register configuration
      address: 40070
      type: holding
      decode: int32
    scale: -1 # reverse direction
```

<a id="meter-e3dc-pv-meter"></a>
#### E3DC (PV Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: e3dc.fritz.box:502
    id: 1 # Modbus Device ID
    register: # manual register configuration
      address: 40067 # (40068 - 1) "Photovoltaikleistung in Watt"
      type: holding
      decode: int32s
    scale: -1 # reverse sign
```

<a id="meter-kostal-inverter-grid-meter"></a>
#### Kostal Inverter (Grid Meter)

```yaml
- type: default
  power:
    type: modbus # use Modbus plugin
    model: kostal
    uri: 192.168.178.52:1502 
    id: 71 # Modbus Device ID 
    register: # manual register configuration
      address: 252 # (see https://www.kostal-solar-electric.com/de-de/download/-/media/document-library-folder---kse/2018/08/30/08/53/ba_kostal_interface_modbus-tcp_sunspec.pdf)
      type: holding
      decode: float32s #swapped float encoding
```

<a id="meter-kostal-inverter-pv-meter"></a>
#### Kostal Inverter (PV Meter)

```yaml
- type: modbus
  model: kostal
  uri: 192.168.0.1:1502
  id: 71 # Modbus Device ID
  power: Power
```

<a id="meter-kostal-smart-energy-meter-grid-meter"></a>
#### Kostal Smart Energy Meter (Grid Meter)

```yaml
- type: modbus
  model: kostal
  uri: 192.168.0.1:502
  id: 71 # Modbus Device ID
  power: Power
  energy: Energy
```

<a id="meter-sma-sunspec-pv-meter"></a>
#### SMA (or other SunSpec compatible) Inverter (PV Meter)

```yaml
- type: modbus
  value: Power
  model: sunspec
  uri: 192.168.178.101:502 # Modbus-TCP address:port
  id: 126 # SMA Modbus Device ID (SunSpec)
```

<a id="meter-multiple-inverters-combined-pv-meter"></a>
#### Multiple SunSpec compatible Inverters combined (PV Meter)

```yaml
- type: default
  power:
    type: calc
    add:
    - type: modbus
      value: Power
      model: sunspec
      uri: 192.168.178.101:502 # Modbus-TCP address:port of Inverter 1
      id: 126 # Modbus Device ID (SunSpec)
    - type: modbus
      value: Power
      model: sunspec
      uri: 192.168.178.102:502 # Modbus-TCP address:port of Inverter 2
      id: 126 # Modbus Device ID (SunSpec)
```

<a id="meter-sma-home-manager-2-0--sma-energy-meter"></a>
#### SMA Home Manager 2.0 / SMA Energy Meter

```yaml
- type: sma
  serial: 1234567890 # Serial number of the device
```

<a id="meter-solarlog-grid-meter"></a>
#### SolarLog (Grid Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: 192.168.0.32:502 # Modbus-TCP address:port
    id: 1 # Modbus Device ID
    register:
      address: 3518
      type: input
      decode: uint32s
```

<a id="meter-solarlog-pv-meter"></a>
#### SolarLog (PV Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: 192.168.0.32:502 # Modbus-TCP address:port
    id: 1 # Modbus Device ID
    register:
      address: 3502
      type: input
      decode: uint32s
```

<a id="meter-vzlogger-http"></a>
#### vzlogger (HTTP)

```yaml
- type: default
  power: # power reading
    type: mqtt # use mqtt plugin
    topic: mbmd/sdm1-1/Power # mqtt topic
    timeout: 10s # don't use older values
```

<a id="meter-vzlogger-push-server-websocket"></a>
#### vzlogger (Push Server/ Websocket)

```yaml
- type: default
  power:
    type: ws # use websocket plugin
    uri: ws://volkszaehler:8082/socket
    jq: .data | select(.uuid=="<uuid>") .tuples[0][1] # parse response json
    timeout: 30s
    scale: 1
```

<a id="meter-vzlogger-split-import-export-channels"></a>
#### vzlogger (split import/export channels)

```yaml
- type: default
  power:
    type: calc # use calc plugin
    add:
    - type: http # import channel
      uri: http://volkszaehler/api/data/<import-uuid>.json?from=now
      jq: .data.tuples[0][1] # parse response json
    - type: http # export channel
      uri: http://volkszaehler/api/data/<export-uuid>.json?from=now
      jq: .data.tuples[0][1] # parse response json
      scale: -1 # export must result in negative values
```

<a id="meter-generic-mqtt"></a>
#### Generic (MQTT)

```yaml
- type: default
  power: # power reading
    type: mqtt # use mqtt plugin
    topic: mbmd/sdm1-1/Power # mqtt topic
    timeout: 10s # don't use older values
```

<a id="meter-generic-script"></a>
#### Generic (Script)

```yaml
- type: default
  power:
    type: script # use script plugin
    cmd: /bin/sh -c "echo 0" # actual command
    timeout: 3s # kill script after 3 seconds
```


### Chargers


<a id="charger-go-echarger-cloud"></a>
#### go-eCharger (Cloud)

```yaml
- type: go-e
  token: 4711c # or go-e cloud API token
  cache: 10s # go-e cloud API cache duration
```

<a id="charger-go-echarger-lokal"></a>
#### go-eCharger (Lokal)

```yaml
- type: go-e
  uri: http://192.168.1.4 # either go-e local address
```

<a id="charger-keba-connect"></a>
#### KEBA Connect

```yaml
- type: keba
  uri: 192.168.1.4:7090 # KEBA address:port
  rfid:
    tag: 765765348 # RFID tag, see `evcc charger` to show tag
```

<a id="charger-mobile-charger-connect"></a>
#### Mobile Charger Connect

```yaml
- type: mcc
  uri: https://192.168.1.4 # Mobile Charger Connect address
  password: # home user password
```

<a id="charger-nrgkick-bluetooth"></a>
#### NRGKick Bluetooth (direct)

```yaml
- type: nrgkick-bluetooth
  macaddress: 00:99:22 # MAC address
  pin: 1234 # pin
```

<a id="charger-nrgkick-connect"></a>
#### NRGKick via Connect device

```yaml
- type: nrgkick-connect
  ip: 192.168.1.4 # IP
  macaddress: 00:99:22 # MAC address
  password: # password
```

<a id="charger-openwb-remote-controlled-using-mqtt"></a>
#### OpenWB (remote-controlled using MQTT)

```yaml
- type: default
  status:
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
    payload: ${maxCurrent:%d} # write payload definition
```

<a id="charger-phoenix-em-cp-controller-modbus-tcp"></a>
#### Phoenix EM-CP Controller (Modbus-TCP)

```yaml
- type: phoenix-emcp
  uri: 192.168.0.8:502 # Modbus-TCP address:port
  id: 1 # Modbus Device ID
```

<a id="charger-phoenix-ev-cc-controller-modbus"></a>
#### Phoenix EV-CC Controller (Modbus)

```yaml
- type: phoenix-evcc
  device: /dev/ttyUSB0
  baudrate: 9600
  comset: 8N1
  id: 1 # Modbus Device ID
```

<a id="charger-simple-evse-tcp"></a>
#### Simple EVSE (Modbus over TCP)

```yaml
- type: simpleevse
  uri: 192.168.0.8:5002 # Modbus-RTU server address:port
```

<a id="charger-simple-evse-usb"></a>
#### Simple EVSE (Modbus on USB)

```yaml
- type: simpleevse
  device: /dev/usb1 # Serial RS485 busmaster device
```

<a id="charger-evse-wifi"></a>
#### EVSE Wifi

```yaml
- type: evsewifi
  uri: http://192.168.1.4 # SimpleEVSE-Wifi IP address
```

<a id="charger-wallbe-eco-pro"></a>
#### Wallbe (Eco, Pro)

```yaml
- type: wallbe
  uri: 192.168.0.8:502 # Modbus-TCP address:port
```

<a id="charger-wallbe-pre-2019-ev-cc-ac1-controller"></a>
#### Wallbe (pre 2019 EV-CC-AC1 controller)

```yaml
- type: wallbe
  uri: 192.168.0.8:502 # Modbus-TCP address:port
  legacy: true # enable for older Wallbes with Phoenix EV-CC-AC1-M3-CBC-RCM controller
```

<a id="charger-generic-mqtt"></a>
#### Generic (MQTT)

```yaml
- type: default
  status: # charger status A..F
    type: mqtt
    topic: some/topic1
  enabled: # charger enabled state (true/false or 0/1)
    type: mqtt
    topic: some/topic2
  enable:
    type: mqtt
    topic: some/topic3
    payload: ${enable:%d} # write payload definition
  maxcurrent:
    type: mqtt
    topic: some/topic4
    payload: ${maxCurrent:%d} # write payload definition
```

<a id="charger-generic-script"></a>
#### Generic (Script)

```yaml
- type: default
  status: # charger status A..F
    type: script
    cmd: /bin/sh -c "echo C"
  enabled: # charger enabled state (true/false or 0/1)
    type: script
    cmd: /bin/sh -c "echo 1"
  enable: # set charger enabled state
    type: script
    cmd: /bin/sh -c "echo ${enable:%d}"
  maxcurrent: # set charger max current
    type: script
    cmd: /bin/sh -c "echo ${maxcurrent:%d}"
```


### Vehicles (State of Charge)


<a id="vehicle-audi"></a>
#### Audi

```yaml
- type: audi
  title: Q55 TFSIe # display name for UI
  capacity: 10 # kWh
  user: # user
  password: # password
  vin: WAUZZZ...
  cache: 5m # cache API response
```

<a id="vehicle-bmw-i3"></a>
#### BMW (i3)

```yaml
- type: bmw
  title: i3 # display name for UI
  capacity: 65 # kWh
  user: # user
  password: # password
  vin: WBMW...
  cache: 5m # cache API response
```

<a id="vehicle-nissan-leaf"></a>
#### Nissan (Leaf)

```yaml
- type: nissan
  title: Leaf # display name for UI
  capacity: 60 # kWh
  user: # user
  password: # password
  region: NE # carwings region, leave empty for Europe
  cache: 5m # cache API response
```

<a id="vehicle-porsche"></a>
#### Porsche

```yaml
- type: porsche
  title: Taycan # display name for UI
  capacity: 83 # kWh
  user: # user
  password: # password
  vin: WP...
  cache: 5m # cache API response
```

<a id="vehicle-renault-zoe"></a>
#### Renault (Zoe)

```yaml
- type: renault
  title: Zoe # display name for UI
  capacity: 40 # kWh
  user: user@mailserver.net # user
  password: safestpasswordever # password
  region: de_DE # gigya region
  vin: WREN...
  cache: 10m # cache API response
```

<a id="vehicle-tesla"></a>
#### Tesla

```yaml
- type: tesla
  title: Model S # display name for UI
  capacity: 90 # kWh
  clientid: # client id
  clientsecret: # client secret
  email: # email
  password: # password
  vin: WTSLA...
  cache: 5m # cache API response
```

<a id="vehicle-generic-script"></a>
#### Generic (Script)

```yaml
- type: default
  title: Auto # display name for UI
  capacity: 50 # kWh
  charge:
    type: script # use script plugin
    cmd: /bin/sh -c "echo 50" # actual command
    timeout: 3s # kill script after 3 seconds
  cache: 5m # cache duration
```

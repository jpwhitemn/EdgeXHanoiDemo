# Hanoi Demo

This demo shows the Hanoi release of EdgeX with the following devices:

- Comet Systems T0310 temperature probe (Modbus RTU)
- Patlite Signal Tower (SNMP)
- Moisture sensor

*The device services - with customizations for these devices - are included with this demo, but see the originally linked repositories for updates to the device service code*

## Starting

To run the demo, perform the following operations:

1. Start EdgeX (Hanoi) using Docker Compose (docker-compose up -d in the ./compose/ folder). This Compose file has been modified for this demo (removal of security protections and non-use device services)
2. Start the device services running on the same host using the start.sh script in this directory
3. Start the moisture device service running on a Raspberry Pi 3


To stop the demo, perform the following operations:
1. Stop the device services with Ctrl-C in the start.sh terminal
2. Stop the moisture device service with Ctrl-C running on the RP3
2. Stope EdgeX using Docker Compose (docker-compose down in the ./compose/ folder)


## Rules Setup

### Create the stream

Send the following POST to the Kuiper Rules Engine
192.168.0.35:48075/streams (with content type set to application/json)

``` JSON
	{
	  "sql": "create stream watches() WITH (FORMAT=\"JSON\", TYPE=\"edgex\")"
	}
```

### Create the rules to monitor the temperature probe and control the Patlite Red Light

**Create the rule to turn on the patlite**

Send the following POST to the Kuiper Rules Engine (replace command IDs as appropriate).

``` JSON

	{
	  "id": "probe-out-of-position",
	  "sql": "SELECT ProbeTemperature FROM watches WHERE ProbeTemperature < 900",
	  "actions": [
	    {
	      "rest": {
	        "url": "http://edgex-core-command:48082/api/v1/device/c4bfef8c-273f-414b-9481-c2ad5e801d15/command/66046772-d3e6-45b4-a267-bf71354a26c3",
	        "method": "put",
	        "retryInterval": -1,
	        "dataTemplate": "{\"RedLightControlState\":\"2\",\"RedLightTimer\":\"0\"}",
	        "sendSingle": true
	      }
	    },
	    {
	      "log":{}
	    }
	  ]
	}

```

**Create the rule to turn off the patlite**

Send the following POST to the Kuiper Rules Engine (replace command IDs as appropriate).

``` JSON
	{
	  "id": "probe-in-position",
	  "sql": "SELECT ProbeTemperature FROM watches WHERE ProbeTemperature >= 900",
	  "actions": [
	    {
	      "rest": {
	        "url": "http://edgex-core-command:48082/api/v1/device/c4bfef8c-273f-414b-9481-c2ad5e801d15/command/66046772-d3e6-45b4-a267-bf71354a26c3",
	        "method": "put",
	        "retryInterval": -1,
	        "dataTemplate": "{\"RedLightControlState\":\"1\",\"RedLightTimer\":\"0\"}",
	        "sendSingle": true
	      }
	    },
	    {
	      "log":{}
	    }
	  ]
	}
```

### Create the rules to monitor the moisture sensor and control the Patlite Amber (yellow) Light

**Create the rule to turn off the patlite**

Send the following POST to the Kuiper Rules Engine (replace command IDs as appropriate).

``` JSON
	{
	"id": "moisture-not-detected",
	"sql": "SELECT MoistureState FROM watches WHERE MoistureState < 1",
	"actions": [
		{
		"rest": {
			"url": "http://edgex-core-command:48082/api/v1/device/c4bfef8c-273f-414b-9481-c2ad5e801d15/command/f5f7b4ad-74ab-4f02-a341-e384646ae250",
			"method": "put",
			"retryInterval": -1,
			"dataTemplate": "{\"AmberLightControlState\":\"1\",\"AmberLightTimer\":\"0\"}",
			"sendSingle": true
		}
		},
		{
		"log":{}
		}
	]
	}
```

**Create the rule to turn on the patlite**

Send the following POST to the Kuiper Rules Engine (replace command IDs as appropriate).

``` JSON
	{
	"id": "moisture-detected",
	"sql": "SELECT MoistureState FROM watches WHERE MoistureState >= 1",
	"actions": [
		{
		"rest": {
			"url": "http://edgex-core-command:48082/api/v1/device/c4bfef8c-273f-414b-9481-c2ad5e801d15/command/f5f7b4ad-74ab-4f02-a341-e384646ae250",
			"method": "put",
			"retryInterval": -1,
			"dataTemplate": "{\"AmberLightControlState\":\"2\",\"AmberLightTimer\":\"0\"}",
			"sendSingle": true
		}
		},
		{
		"log":{}
		}
	]
	}
```
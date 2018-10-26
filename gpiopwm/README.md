# GPIO PWM activity
This activity allows you control PWM pins of a Raspberry Pi.

> **WARNING: The GPIO PWM activity requires that flogo runs as root.**

## Installation
### Flogo Web

#### Start

Start a container of Flogo Web UI :

```bash
docker run --name flogo -it -d -p 3303:3303 -e FLOGO_NO_ENGINE_RECREATION=false flogo/flogo-docker eula-accept
```
*The environment variable FLOGO_NO_ENGINE_RECREATION=false allows to force import of installed contributions.*

#### Installation of the activity

To install the activity into the started container :

```bash
docker exec -it flogo sh -c 'cd /tmp/flogo-web/build/server/local/engines/flogo-web && flogo install github.com/square-it/flogo-contrib-activities/gpiopwm'
```

Restart the container
```bash
docker restart flogo
```

### Flogo CLI
```bash
flogo install github.com/square-it/flogo-contrib-activities/gpiopwm
```

## Schema
Inputs and Outputs:

```json
{
 "inputs":[
    {
      "name": "pinNumber",
      "type": "int",
      "required": true,
      "allowed": [12, 13, 18, 19, 40, 41, 45]
    },
    {
      "name": "pwmFrequency",
      "type": "int",
      "required": true
    },
    {
      "name": "dutyLength",
      "type": "uint32",
      "required": true
    },
    {
      "name": "cycleLength",
      "type": "uint32",
      "required": true
    }
  ],
  "outputs": [
  ]
 }
```
## Settings
| Setting      | Required | Description |
|:-------------|:---------|:------------|
| pinNumber    | True     | The index of the pin of the GPIO header to modify |
| pwmFrequency | True     | The frequency (in Hz) of the PWM |
| dutyLength   | True     | The duty part of the duty cycle |
| cycleLength  | True     | The cycle part of the duty cycle |

See https://godoc.org/github.com/stianeikeland/go-rpio#SetDutyCycle for more details about Duty and Cycle.

## Outputs

| Output     | Description |
|:------------|:---------|

## Examples
### Activate PWM with 1ms duty and 50Hz frequency on pin 12

```json
{
  "id": "gpiopwm_1",
  "name": "GPIO PWM",
  "description": "Perform PWM control on Raspberry GPIO. Requires root access.",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/command",
    "input": {
      "pinNumber": 12,
      "pwmFrequency": 50,
      "dutyLength": 100,
      "cycleLength": 2000
    },
    "output": {
    }
  }
}
```

### Disable PWM on pin 12 (set dutyLength to 0)

```json
{
  "id": "gpiopwm_1",
  "name": "GPIO PWM",
  "description": "Perform PWM control on Raspberry GPIO. Requires root access.",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/command",
    "input": {
      "pinNumber": 12,
      "pwmFrequency": 50,
      "dutyLength": 0,
      "cycleLength": 2000
    },
    "output": {
    }
  }
}
```

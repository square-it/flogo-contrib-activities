# Sleep activity

This activity allows you to make the current thread sleep for a specified amount of time.

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
docker exec -it flogo sh -c 'cd /tmp/flogo-web/build/server/local/engines/flogo-web && flogo install github.com/square-it/flogo-contrib-activities/sleep'
```

Restart the container
```bash
docker restart flogo
```

### Flogo CLI
```bash
flogo install github.com/square-it/flogo-contrib-activities/sleep
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "duration",
      "type": "string",
      "required": "true"
    }
  ],
  "outputs": [

  ]
}
```

## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| duration    | True     | The amount of a time during which the thread will be sleeping.<br />This duration can be expressed by any valid expression described by [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration) Go method.<br />For instance: "500ms", "5s", "1m30s" |

## Outputs

No output

## Examples

### Simple example

```json
{
  "id": "sleep_1",
  "name": "Sleep5s",
  "description": "Sleeps for five seconds",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/sleep",
    "input": {
      "duration": "5s"
    },
    "output": {
    }
  }
}
```


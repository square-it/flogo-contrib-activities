# Copy File
This activity allows you to copy a file.

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
docker exec -it flogo sh -c 'cd /tmp/flogo-web/build/server/local/engines/flogo-web && flogo install github.com/square-it/flogo-contrib-activities/copyfile'
```

Restart the container
```bash
docker restart flogo
```

### Flogo CLI
```bash
flogo install github.com/square-it/flogo-contrib-activities/copyfile
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "source",
      "type": "string",
      "required": true
    },
    {
      "name":"destination",
      "type":"string",
      "required": true
    }
  ],
  "outputs": [
  ]
 }
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| source      | True     | The source file |         
| destination | True     | The destination file |


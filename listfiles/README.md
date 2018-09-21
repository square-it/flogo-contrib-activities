# List Files
This activity allows you to list filenames of a directory recursively or not.

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
docker exec -it flogo sh -c 'cd /tmp/flogo-web/build/server/local/engines/flogo-web && flogo install github.com/square-it/flogo-contrib-activities/listfiles'
```

Restart the container
```bash
docker restart flogo
```

### Flogo CLI
```bash
flogo install github.com/square-it/flogo-contrib-activities/listfiles
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "directory",
      "type": "string"
    },
    {
      "name":"recursive",
      "type":"boolean"
    }
  ],
  "outputs": [
    {
      "name": "filenames",
      "type": "array"
    }
  ]
 }
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| directory   | False    | The directory to be listed |         
| recursive   | False    | If this field is set to true, list recursively. This field defaults to false |
| filenames   | False    | The list of filenames |


## Examples
### List

```json
{
  "id": "listfiles_1",
  "name": "List files",
  "description": "List files activity",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/listfiles",
    "input": {
      "directory": "/tmp",
      "recursive": false
    }
  }
}
```

### List recursively

```json
{
  "id": "listfiles_1",
  "name": "List files",
  "description": "List files activity",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/listfiles",
    "input": {
      "directory": "/tmp",
      "recursive": true
    }
  }
}
```

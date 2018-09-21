# Remove File
This activity allows you to remove a file or a directory. With the option recursive, it is possible to remove a directory and his content.

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
docker exec -it flogo sh -c 'cd /tmp/flogo-web/build/server/local/engines/flogo-web && flogo install github.com/square-it/flogo-contrib-activities/removefile'
```

Restart the container
```bash
docker restart flogo
```

### Flogo CLI
```bash
flogo install github.com/square-it/flogo-contrib-activities/removefile
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "path",
      "type": "string"
    },
    {
      "name":"recursive",
      "type":"boolean",
      "value": false
    }
  ],
  "outputs": [
  ]
 }
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| directory   | False    | The path to be removed |         
| recursive   | False    | If this field is set to true, remove recursively a directory. It has no impact for a file. This field defaults to false. |


## Examples
### Remove a file

```json
{
  "id": "removefile_1",
  "name": "Remove File",
  "description": "Remove a file or an directory",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/removefile",
    "input": {
      "directory": "/tmp/file",
      "recursive": false
    }
  }
}
```

### Remove a directory recursively

```json
{
  "id": "removefile_1",
  "name": "Remove File",
  "description": "Remove a file or an directory",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/removefile",
    "input": {
      "directory": "/tmp/dir",
      "recursive": true
    }
  }
}
```

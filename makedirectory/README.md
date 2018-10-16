# Make Directory
This activity allows you to make a directory. With the options, it is possible to create all parents directories if necessary and set the permissions.

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
docker exec -it flogo sh -c 'cd /tmp/flogo-web/build/server/local/engines/flogo-web && flogo install github.com/square-it/flogo-contrib-activities/makedirectory'
```

Restart the container
```bash
docker restart flogo
```

### Flogo CLI
```bash
flogo install github.com/square-it/flogo-contrib-activities/makedirectory
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "path",
      "type": "string",
      "required": "true"
    },
    {
      "name":"all",
      "type":"boolean",
      "value": false
    },
    {
      "name":"permissions",
      "type":"string",
      "value": "0777"
    }
  ],
  "outputs": [
  ]
 }
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| path        | True     | The path to be created |         
| all         | False    | If this field is set to true, create all parents directories if necessary. This field defaults to false. |
| permissions | False    | Set the permissions of the directory or directorys if all option is set to true. This field defaults to 0777. The format of the permissions is Unix permission bits. |


## Examples
### Make a directory

This example make the directory */tmp/dir*.

```json
{
  "id": "makedirectory_1",
  "name": "Make Directory",
  "description": "Make a directory",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/makedirectory",
    "input": {
      "path": "/tmp/dir",
    }
  }
}
```

### Make a directory and all these parents

This example make the directories */tmp/dir1*, */tmp/dir1/dir2*, */tmp/dir1/dir2/dir3*.

```json
{
  "id": "makedirectory_1",
  "name": "Make Directory",
  "description": "Make a directory",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/makedirectory",
    "input": {
      "path": "/tmp/dir1/dir2/dir3",
      "all": true
    }
  }
}

```

### Make a directory with permissions

This example make the directory */tmp/dir* with permissions 0700 or drwx------.

```json
{
  "id": "makedirectory_1",
  "name": "Make Directory",
  "description": "Make a directory",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/makedirectory",
    "input": {
      "path": "/tmp/dir",
      "permissions": "0700"
    }
  }
}
```


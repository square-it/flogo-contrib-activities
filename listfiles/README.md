# Counter
This activity allows you to list filenames of a directory recursively or not.

## Installation
### Flogo Web
This activity is not available with the Flogo Web UI
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
| recursive   | False    | If this field is set to true, list recursively. |
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

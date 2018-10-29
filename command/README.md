# Command activity
This activity allows you execute a command.

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
docker exec -it flogo sh -c 'cd /tmp/flogo-web/build/server/local/engines/flogo-web && flogo install github.com/square-it/flogo-contrib-activities/command'
```

Restart the container
```bash
docker restart flogo
```

### Flogo CLI
```bash
flogo install github.com/square-it/flogo-contrib-activities/command
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "command",
      "type": "string",
      "required" : "true"
    },
    {
      "name": "arguments",
      "type": "array"
    },
    {
      "name": "directory",
      "type": "string"
    },
    {
      "name": "environment",
      "type": "array"
    },
    {
      "name":"useCurrentEnvironment",
      "type":"boolean",
      "value": true
    },
    {
      "name":"timeout",
      "type":"integer",
      "value": 0
    },
    {
      "name":"wait",
      "type":"boolean",
      "value": true
    }
  ],
  "outputs": [
    {
      "name": "output",
      "type": "string"
    },
    {
      "name": "exitCode",
      "type": "integer"
     },
    {
      "name":"success",
      "type":"boolean"
    }
  ]
 }
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| command     | True     | The command to be executed. it can be relative or absolute. if it is relative, the command is looked up in the PATH of the operating systeme. |
| arguments   | False    | The arguments of the command. |
| directory   | False    | The working directory of the command. Default value is the empty string, in this case execute the command in the current directory. |
| environment | False    | The environment variables. The format of each element is "\<name\>=\<value\>".
| useCurrentEnvironment | False | True to use the current environment variable. Default value : true. |    
| timeout     | False    | The command is canceled if the number of timeout setting in seconds is exceeded. If this value is equals to 0, the command is never canceled. Default value : 0 |
| wait        | False    | If wait is equals to true, wait the end of exection of the command, else don't wait the end of execution. Default value : true.

## Outputs

| Output     | Description |
|:------------|:---------|
| output      | The output of the command execution. |
| exitCode    | The exit code of the command exection. This value depends on the operating system. In case the operating system is not supported by the activity, the value will be 0 in case of success, -100 otherwise. |
| success     | True if the command execution is successful. False otherwise. |

## Examples
### Simple example

```json
{
  "id": "command_1",
  "name": "Command ls",
  "description": "List a directory",
  "activity": {
    "ref": "github.com/square-it/flogo-contrib-activities/command",
    "input": {
      "command": "ls"
    },
    "output": {
    }
  }
}
```


# wait-for-it

`wait-for-it` is a command line tool that wait for services healthchecks before executing an entrypoint command.  
This tool prevents the application from crash loop until interdependent services are available.  

## Quick Start

You can get `wait-for-it` command line tool from the [releases](https://github.com/efi1397/wait-for-it/releases) section in this `github` project.   

```console
foo@bar:~$ wait-for-it

Usage:
  wait-for-it [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  wait        Wait for hosts healthcheks

Flags:
  -h, --help     help for wait-for-it
  -t, --toggle   Help message for toggle

Use "wait-for-it [command] --help" for more information about a command.
```

Suppose we want to execute an entrypoint which depends on database host (`http://database/healthchecks`) and message broker host (`http://rabbitmq/healthchecks`).  
For this scenario we need to execute this command:  
```console
foo@bar:~$ wait-for-it wait --hosts  http://rabbitmq/healthchecks,http://database/healthchecks --entrypoint "echo AWESOME!" --timeout 120
```



For more examples and details look at the [examples folder](../examples/).







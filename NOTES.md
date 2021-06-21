# Notes on how to use this service

# Commandline app

## Building

This requires a `go1.16+` compiler. If you have cloned this repo then build it by running.

```shell
mkdir bin
make
```

This will create `bin/voting_service` binary.

## Commandline options

```shell
NAME:
   Voting Service - Simple gRPC service for voting

USAGE:
   voding_service [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value      Port to bind the service (default: 8080) [$VOTING_SERVICE_PORT]
   --config value, -c value    Path to the configuration file [$VOTING_CONFIG_FILE]
   --region value, -r value    aws region (default: "local") [$VOTING_AWS_REGION]
   --endpoint value, -e value  dynamodb endpoint (default: "http://localhost:8000") [$VOTING_DYNAMODB_ENDPOINT]
   --help, -h                  show help
```

# Starting the service

```shell
./bin/voding_service
```
If all is well you should see logs resembling these

```shell
$ ./bin/voding_service
{"level":"info","ts":1624314016.975709,"logger":"main","msg":"Starting listener for votable service","port":0}
{"level":"info","ts":1624314016.982441,"logger":"main","msg":"Opening storage"}
{"level":"info","ts":1624314016.982466,"logger":"main","msg":"Setting up open telemetry"}
```
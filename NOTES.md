

# Notes on how to use this service

# Completed features

| Functionality | Status |
|----------------|---------|
| Go gRPC service | ✅ |
| Dynamo based store | ✅ |
| Pagination for `ListVoteables` RPC call | ✅ |
| Open telemetry tracing with stdout export support | ✅ |
| Prometheus metrics | ❌️ |
| Configuration and secret management | ✅ |


# Limitation

- Only one vote cast per voteable

# Commandline app

## Building

This requires a `go1.16+` compiler. If you have cloned this repo then build it by running.

```shell
mkdir -p bin
make
```

This will create `bin/voting_service` binary.

## Commandline options

```shell
NAME:
   Voting Service - Simple gRPC service for voting

USAGE:
   voting_service [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value      Port to bind the service (default: 8080) [$VOTING_SERVICE_PORT]
   --config value, -c value    Path to the configuration file [$VOTING_CONFIG_FILE]
   --region value, -r value    aws region (default: "local") [$VOTING_AWS_REGION]
   --endpoint value, -e value  dynamodb endpoint (default: "http://localhost:8000") [$VOTING_DYNAMODB_ENDPOINT]
   --mem, -m                   Uses an in memory storage [$VOTING_MEMORY_STORE]
   --trace, -t                 Enable open tracing, the traces will be exported to stdout [$VOTING_TRACE]
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

# How to use

There is a client sample included in `cmd/voting_client/` that connects to the service and perform simple tasks


## CreateVoteable

```shell
--> request
{"level":"info","ts":1624352262.817427,"logger":"main.Client","msg":"client request payload logged as grpc.request.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"CreateVoteable","grpc.request.content":{"msg":{"question":"0 - Truth or Dare?","answers":["Truth","Dare"]}}}
--> response
{"level":"info","ts":1624352262.831533,"logger":"main.Client","msg":"client response payload logged as grpc.response.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"CreateVoteable","grpc.response.content":{"msg":{"uuid":"7e97e076-f1ec-4620-93d4-81124cb78d3a"}}}
```

## ListVoteables all

```shell
--> request
{"level":"info","ts":1624353109.160086,"logger":"main.Client","msg":"client request payload logged as grpc.request.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"ListVoteables","grpc.request.content":{"msg":{}}}

--> response
{"level":"info","ts":1624353109.1704152,"logger":"main.Client","msg":"client response payload logged as grpc.response.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"ListVoteables","grpc.response.content":{"msg":{"votables":[{"uuid":"fd97451b-c613-473b-ac20-28678d3b3c68","question":"1 - Truth or Dare?","answers":["Dare","Truth"]},{"uuid":"bad655c3-9408-416f-b375-05a31b799283","question":"2 - Truth or Dare?","answers":["Dare","Truth"]},{"uuid":"787555e1-6105-4525-8540-74fa56c16d6f","question":"0 - Truth or Dare?","answers":["Dare","Truth"]}],"lastIndex":"e30="}}}
```

## ListVoteables  list with limit

```shell
--> request
{"level":"info","ts":1624353109.17052,"logger":"main.Client","msg":"client request payload logged as grpc.request.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"ListVoteables","grpc.request.content":{"msg":{"limit":2}}}

--> response
{"level":"info","ts":1624353109.1789498,"logger":"main.Client","msg":"client response payload logged as grpc.response.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"ListVoteables","grpc.response.content":{"msg":{"votables":[{"uuid":"fd97451b-c613-473b-ac20-28678d3b3c68","question":"1 - Truth or Dare?","answers":["Dare","Truth"]},{"uuid":"bad655c3-9408-416f-b375-05a31b799283","question":"2 - Truth or Dare?","answers":["Dare","Truth"]}],"lastIndex":"eyJ1dWlkIjoiYmFkNjU1YzMtOTQwOC00MTZmLWIzNzUtMDVhMzFiNzk5MjgzIn0="}}}
```

## ListVoteables  list with pagination

```shell
--> request
{"level":"info","ts":1624353109.179117,"logger":"main.Client","msg":"client request payload logged as grpc.request.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"ListVoteables","grpc.request.content":{"msg":{"lastIndex":"eyJ1dWlkIjoiYmFkNjU1YzMtOTQwOC00MTZmLWIzNzUtMDVhMzFiNzk5MjgzIn0="}}}

--> response
{"level":"info","ts":1624353109.188828,"logger":"main.Client","msg":"client response payload logged as grpc.response.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"ListVoteables","grpc.response.content":{"msg":{"votables":[{"uuid":"787555e1-6105-4525-8540-74fa56c16d6f","question":"0 - Truth or Dare?","answers":["Dare","Truth"]}],"lastIndex":"e30="}}}
```

## CastVote

```shell
--> request
{"level":"info","ts":1624355156.177943,"logger":"main.Client","msg":"client request payload logged as grpc.request.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"CastVote","grpc.request.content":{"msg":{"uuid":"78c7699b-37e9-425a-a477-f068cdc143f2"}}}

--> response
{"level":"info","ts":1624355156.347154,"logger":"main.Client","msg":"client response payload logged as grpc.response.content","system":"grpc","span.kind":"client","grpc.service":"api.VotingService","grpc.method":"CastVote","grpc.response.content":{"msg":{"status":"ok"}}}
```
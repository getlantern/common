This repository contains common Go data structures and algorithms shared across Lantern Go projects.

One such struct is ChainedServerInfo, which can be generated from the associated protocol buffers
definition file with the following:

## Prerequisites

- protoc (brew install protobuf)
- Go protobufs (go install google.golang.org/protobuf/cmd/protoc-gen-go@latest)

With the above prerequisites installed, you can run:

```
protoc --go_out=paths=source_relative:. server.proto
```
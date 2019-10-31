# NAVILLERA

> Navillera is a service boilerplate using golang and grpc connection running in docker container

## Prerequisite

- Install [go](https://golang.org/doc/install) in local machine for convenience development experience (auto complete, code sugestion, etc)
- Install golang plugin to your editor choice (ie. VSCode, Atom)
- [Docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/)
- Protoc and proto-gen-go plugin (described in the next session)

## Protoc and Proto-gen-go

### Install protoc and protoc-gen-go in Mac

- Install protoc-gen-go use homebrew `brew install protoc-gen-go` it will also install protoc

### Install protoc and protoc-gen-go in Ubuntu

- Install protoc-gen-go `go get -u github.com/golang/protobuf/protoc-gen-go`. It will also install protoc
- The compiler plugin, protoc-gen-go, will be installed in `$GOPATH/bin` unless `$GOBIN` is set. It must be in your $PATH for the protocol compiler, protoc, to find it

### Install protoc and protoc-gen-go in Windows

- TODO

## How to run

Despite, it is possible to run this project in local machine we strongly recommend you to run it using docker. Please follow this steps:

- generate protobuf file if you have not generated it yet or you made changes to proto file `protoc -I proto/ proto/call.proto --go_out=plugins=grpc:proto`
- using docker-compose run `docker-compose -f docker-compose.dev.yaml`
- navillera is available in `localhost:50050` you can access it using RPC client your choices

## Recommended rpc client

- [bloomRPC](https://github.com/uw-labs/bloomrpc)

## Nice to read

- [Golang playground](https://play.golang.org/)
- [Go module](https://blog.golang.org/using-go-modules)
- [gRPC](https://grpc.io/docs/guides/concepts/)
- [Golang mongo driver](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
- [Docker for mac](https://docs.docker.com/docker-for-mac/)
- [Docker compose](https://docs.docker.com/compose/)

## TODO

- Change procedure entity placeholder with real world scenario included but not limited to `get`, `getById`, `post`, `put` and `delete`
- Add mongo store interface
- Guide to install protoc and protoc-gen-go in Windows

# A Service to manage Test's state, from Software Capstone Project, using gRPC API

## How to compile the gRPC Proto for Go
Compile the proto files with below command: <br />
`protoc ./protos/teststatemanagerpb.proto \` <br />
    `--go-grpc_out=./src --go-grpc_opt=paths=source_relative \` <br />
    `--go_out=./src --go_opt=paths=source_relative`


## How to run the code
From the root directory, run the below command: <br />
`go run ./src/main.go`
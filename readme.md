# A Service to manage Test's state, from Software Capstone Project, using gRPC API

## How to compile the gRPC Proto for Go (No need to follow this section for deployment)
Compile the proto files with below command: <br />
`protoc ./protos/teststatemanagerpb.proto \` <br />
    `--go-grpc_out=./src --go-grpc_opt=paths=source_relative \` <br />
    `--go_out=./src --go_opt=paths=source_relative`


## How to build the code
### Requirement
- go version go1.18.1
- .env (secret)

### Build and run the code

1. From the root directory, run the below command to compile the code into the executable: <br />
`go build ./src/main.go`

2. After compiling, we could run the executable code by using the below command: 
`./main`

## API documentations

### Start Test
#### Request

```javascript
message StartTestRequest {
    int32 userId = 1;
    int32 testId = 2;
}
```

#### Response

```javascript
message StartTestResponse {
    int32 testAttemptId = 1;
}
```

### Get Inprogress Test
#### Request
```javascript
message GetInprogressTestRequest {
    int32 testAttemptId = 1;
}
```

#### Response
```javascript
message GetInprogressTestResponse {
    int32 timeRemainInSecond = 1;
}
```

### End Test (update regularly)
#### Request
```javascript
message EndTestRequest {
    int32 testAttemptId = 1;
}
```

#### Response
```javascript
message EndTestResponse {
    int32 errorCode = 1;
    string message = 2;
}
```
FROM golang:1.18.1

WORKDIR /TestStateManager

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 8060

COPY .env /

COPY src/ ./src/

RUN CGO_ENABLED=0 GOOS=linux go build ./src/main.go

CMD ["./main"]

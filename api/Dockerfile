FROM golang:1.14

WORKDIR source

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# TODO temporary until the fabric is created
env DATA_SERVICE_HOST localhost
env DATA_SERVICE_PORT 8080
env REDIS_HOST localhost
env REDIS_PORT 6379
env SERVER_PORT 8081

CMD ["go", "run", "."]

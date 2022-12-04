FROM golang:1.19.3-bullseye


WORKDIR /app
CMD ["app"]

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/main/main.go

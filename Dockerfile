FROM golang:1.18

WORKDIR /go/src

COPY go.mod .
COPY go.sum .

COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]

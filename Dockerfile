FROM golang:1.26.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o obsys ./cmd/server/

EXPOSE 8090

CMD ["./obsys"]
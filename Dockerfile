FROM golang:1.22

WORKDIR /app

COPY  go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o app ./cmd/main.go

EXPOSE 8080

CMD ["./app"]

FROM golang:1.22


WORKDIR /app

COPY  go.mod go.sum ./
RUN go mod download
RUN apt-get update && apt-get install -y netcat-traditional

COPY . ./
COPY .env.example .env
RUN go build -o app ./cmd/main.go
# ENV GIN_MODE=release

EXPOSE 8080

CMD ["./start.sh"]

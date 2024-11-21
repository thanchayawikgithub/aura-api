FROM golang:1.22.6

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o aura cmd/aura/main.go

EXPOSE 8080

CMD ["./aura"]

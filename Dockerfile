FROM golang:1.21.4

WORKDIR /app

COPY . .

RUN go build -o auth-server .

EXPOSE 8080

CMD ["./auth-server"]

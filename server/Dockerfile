FROM golang:1.18

WORKDIR /go/src/app

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main"]
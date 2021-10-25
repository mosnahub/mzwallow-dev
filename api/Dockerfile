FROM golang:1.17

WORKDIR /go/src/app

COPY go.mod .
RUN go mod tidy

COPY . .

RUN go build -o main .

EXPOSE 5000

CMD ["./main"]
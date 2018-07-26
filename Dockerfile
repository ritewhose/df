FROM golang:1.9

WORKDIR /go/src/gtb7

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep && dep ensure

RUN go build -o /app/gtb7 ./cmd/bot

WORKDIR /app

CMD ["./gtb7"]

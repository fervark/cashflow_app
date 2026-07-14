FROM golang:1.26.5 AS builder

WORKDIR /var/www/cashflow_app

COPY . .

RUN go mod tidy
RUN go build -o main "./cmd/app/main.go"

CMD [". /main"]

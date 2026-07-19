FROM golang:1.26.5 AS builder

WORKDIR /var/www/cashflow_app

ARG DOCKER_USER=default_user

COPY . .

RUN useradd -m -s /bin/bash -G sudo docker_user
RUN usermod -aG root docker_user

RUN go mod tidy
RUN go build -o main "./cmd/app/main.go"

CMD [". /main"]

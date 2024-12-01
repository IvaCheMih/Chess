FROM golang:1.23.2 AS builder

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download

COPY ./src/ ./src

RUN go build -o chess ./src

FROM ubuntu

WORKDIR /app

ADD ./migrations/ ./migrations/
ADD .env .env

COPY --from=builder /app/chess .

EXPOSE 8080

CMD [ "./chess" ]

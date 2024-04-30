FROM golang:1.21.4 AS builder

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download

COPY ./server/ ./server

RUN go build -o chess ./server

FROM ubuntu

WORKDIR /app

ADD ./migrations/ ./migrations/

COPY --from=builder /app/chess .

EXPOSE 8080

CMD [ "./chess" ]

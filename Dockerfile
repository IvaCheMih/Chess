FROM golang:1.21.4 AS builder

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download

COPY ./server/ ./server

RUN go build -o chess_http ./src

FROM ubuntu

WORKDIR /app

ADD ./postgresql/ ./postgresql/

ADD ./migrations/ ./migrations/

COPY --from=builder /app/chess_http .

EXPOSE 8082

CMD [ "./chess_http" ]

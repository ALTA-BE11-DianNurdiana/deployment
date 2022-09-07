FROM golang:1.18

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go build -o be11-api

CMD ["./be11-api"]
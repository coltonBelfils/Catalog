FROM golang:1.15.11-alpine3.12

WORKDIR /app/

COPY . .

RUN go get

RUN go build -o /app/bin/catalog

EXPOSE 80

ENTRYPOINT ["/app/bin/catalog"]
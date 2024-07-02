FROM golang:1.22

WORKDIR /app

COPY ./datetimehttp/go.mod ./datetimehttp/go.sum ./

RUN go mod download

COPY ./datetimehttp/ .

ENV PORT=8080

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]

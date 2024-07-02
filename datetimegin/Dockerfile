FROM golang:1.22

WORKDIR /app

COPY ./datetimegin/go.mod ./datetimegin/go.sum ./

RUN go mod download

COPY ./datetimegin/ .

ENV PORT=8080

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]

FROM golang

WORKDIR /app

COPY . .
RUN go mod download

RUN go build src/cmd/apiserver/main.go

EXPOSE 8080

CMD ["./main"]
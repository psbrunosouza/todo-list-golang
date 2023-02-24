FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build ./...

EXPOSE 1323

CMD ["/server"]




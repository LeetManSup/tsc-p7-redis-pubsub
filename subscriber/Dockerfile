FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./subscriber .

RUN go build -o subscriber .

CMD ["./subscriber"]

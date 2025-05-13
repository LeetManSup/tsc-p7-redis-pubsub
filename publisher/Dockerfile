FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./publisher .

RUN go build -o publisher .

CMD ["./publisher"]

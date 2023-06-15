FROM golang:1.20-alpine3.17
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /usr/bin/bank-server .
CMD [ "bank-server" ]

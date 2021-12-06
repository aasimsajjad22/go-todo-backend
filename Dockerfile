FROM golang:1.13-alpine

ENV WORKDIR=/app
WORKDIR $WORKDIR

COPY src/go.mod ./
COPY src/go.sum ./

RUN go mod download
COPY src $WORKDIR

RUN go build -o go-todo .
EXPOSE 8080
CMD ["./go-todo"]
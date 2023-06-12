FROM golang:1.19-alpine AS development

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
RUN air init

COPY . .

EXPOSE 3000

CMD air
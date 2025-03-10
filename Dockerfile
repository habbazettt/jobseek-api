FROM golang:1.24.0-alpine3.20

WORKDIR /app

RUN apk add --no-cache git curl

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

ENV PATH="/go/bin:$PATH"

RUN go install github.com/air-verse/air@latest

COPY . .

EXPOSE 8080

CMD ["air"]
FROM golang:1.20-alpine3.17

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o tboard ./cmd/tBoard/main.go
EXPOSE 8089
CMD [ "./tboard" ]
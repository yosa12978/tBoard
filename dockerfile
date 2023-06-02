FROM golang:1.20-alpine3.17

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o tboard ./cmd/tBoard/main.go
EXPOSE 8089
CMD [ "./tboard" ]
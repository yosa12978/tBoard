MAIN_FILE = ./cmd/tBoard/main.go
OUT_FILE = tBoard

run:
	go run ${MAIN_FILE}

build:
	go mod tidy
	go build -o ${OUT_FILE} ${MAIN_FILE}

deps:
	go mod tidy
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go get -u github.com/gin-gonic/gin
	go get -u github.com/joho/godotenv
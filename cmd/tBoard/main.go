package main

import (
	"github.com/joho/godotenv"
	"github.com/yosa12978/tBoard/internal/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	app.Run()
}

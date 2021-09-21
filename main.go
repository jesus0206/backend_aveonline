package main

import (
	// "github.com/joho/godotenv"
	"github.com/joho/godotenv"
	"jesus.tn79/aveonline/app"
)

func init() {
	godotenv.Load()
}

func main() {
	app.StartApp()
}

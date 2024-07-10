package main

import (
	"github.com/outagelab/outagelab/internal/app"
)

func main() {
	olapp := app.New()
	olapp.Start()
}

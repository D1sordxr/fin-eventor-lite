package main

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/bootstrap"
)

func main() {
	app := bootstrap.NewApp()
	app.Run()
}

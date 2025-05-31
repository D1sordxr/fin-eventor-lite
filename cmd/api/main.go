package main

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/bootstrap/api"
)

func main() {
	app := api.NewApp()
	app.Run()
}

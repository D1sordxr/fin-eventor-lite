package main

import "github.com/D1sordxr/fin-eventor-lite/internal/bootstrap/worker"

func main() {
	app := worker.NewApp()
	app.Run()
}

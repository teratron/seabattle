package main

import "github.com/teratron/seabattle/cmd/seabattle/app"

func main() {
	// New создаём новое приложение
	a := app.New()

	// Server создаём сервер приложения и запускаем его
	a.Server()

	// Run запускает приложение
	a.Run()
}

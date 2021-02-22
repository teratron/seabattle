package main

import "github.com/teratron/seabattle/pkg/app"

func main() {
	// New создаёт новое приложение
	a := app.New()

	// Server создаёт сервер приложения и запускаем его
	a.Server()

	// Run запускает приложение
	a.Run()
}

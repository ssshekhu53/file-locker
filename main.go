package main

import (
	"runtime"

	"gofr.dev/pkg/gofr"

	"ssshekhu53/file-locker/constants"
	"ssshekhu53/file-locker/handlers"
	"ssshekhu53/file-locker/services"
	cryptPkg "ssshekhu53/file-locker/services/crypt"
	"ssshekhu53/file-locker/services/unix"
)

func main() {
	var service services.FileLocker

	app := gofr.NewCMD()

	crypt, err := cryptPkg.New()
	if err != nil {
		app.Logger().Fatalf("Error occurred: %v", err)
	}

	switch runtime.GOOS {
	case constants.Darwin, constants.Linux:
		service = unix.New(crypt)
	default:
		app.Logger().Fatalf("Unsupported architecture: %v", runtime.GOOS)
	}

	handler := handlers.New(service)

	app.SubCommand("init", handler.Init)
	app.SubCommand("unlock", handler.Unlock)
	app.SubCommand("lock", handler.Lock)
	app.SubCommand("help", handler.Help)

	app.Run()
}

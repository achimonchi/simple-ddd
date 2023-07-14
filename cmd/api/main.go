package main

import (
	"runtime"

	"github.com/gofiber/fiber/v2"

	"simple-ddd/domain/auth"
)

func main() {
	// set running with total CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: "Simple DDD by NoobeeID",
	})

	api := router.Group("api")
	auth.Run(api)

	router.Listen(":8181")
}

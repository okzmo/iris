package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/okzmo/noname/grpc"
	"github.com/okzmo/noname/util"
	"github.com/okzmo/noname/websocket"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	websocket.SetupWebSocketRoutes(app)

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Set-Cookie",
		AllowCredentials: true,
	}))

	go grpc.RunGRPCGateway(config)
	grpc.RunGRPCServer(config)
}

package main

import (
	"print-requests/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

var (
	version = "development"
)

func main() {
	// init logger
	logger.InitLogger(version)

	r := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	r.Get("/*", index)

	logger.Logger.Info().Msg("app listening on 0.0.0.0:8080")
	r.Listen(":8080")
}

func index(ctx *fiber.Ctx) error {
	d := make(fiber.Map)
	d["hostname"] = ctx.Hostname()
	d["path"] = ctx.Path()
	d["headers"] = string(ctx.Request().Header.Header())
	d["body"] = string(ctx.Body())

	logger.Logger.Info().Interface("request", d).Msg("")
	return ctx.Status(200).JSON(d)
}

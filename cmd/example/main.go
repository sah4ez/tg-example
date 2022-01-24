package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/sah4ez/tg-example/pkg/adder"
	"github.com/sah4ez/tg-example/pkg/transport"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	log.Log().Msg("hello world")
	defer log.Log().Msg("good bye")

	svcAdder := adder.New()

	services := []transport.Option{
		transport.Adder(transport.NewAdder(log.Logger, svcAdder)),
	}

	srv := transport.New(log.Logger, services...).WithLog(log.Logger)

	srv.Fiber().Get("/api/healthcheck", func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusOK)
		return ctx.JSON(map[string]string{"status": "Ok"})
	})

	go func() {
		log.Info().Str("bind", ":9000").Msg("listen on") // Move to config
		if err := srv.Fiber().Listen(":9000"); err != nil {
			log.Panic().Err(err).Stack().Msg("server error")
		}
	}()

	<-shutdown
}

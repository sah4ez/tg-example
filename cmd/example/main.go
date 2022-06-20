package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/sah4ez/tg-example/pkg/config"
	"github.com/sah4ez/tg-example/pkg/errors"
	"github.com/sah4ez/tg-example/pkg/files"
	"github.com/sah4ez/tg-example/pkg/storage"
	"github.com/sah4ez/tg-example/pkg/transport"
	"github.com/sah4ez/tg-example/pkg/user"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	log.Logger = config.Service().Logger()

	log.Log().Msg("hello world")
	defer log.Log().Msg("good bye")

	var err error

	store, err := storage.New(config.Service().DSN)
	errors.ExitOnError(err, "create storage")

	err = store.Migrate()
	errors.ExitOnError(err, "failed migrate")

	var (
		userStore storage.User = store
	)

	svcUser := user.New(userStore)
	svcFiles := files.New()

	services := []transport.Option{
		transport.User(transport.NewUser(svcUser)),
		transport.Files(transport.NewFiles(svcFiles)),
	}

	srv := transport.New(log.Logger, services...).WithLog()

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

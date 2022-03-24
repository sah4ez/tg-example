package errors

import (
	"os"

	"github.com/rs/zerolog/log"
)

func ExitOnError(err error, msg string) {
	if err != nil {
		log.Err(err).Msg(msg)
		os.Exit(1)
	}
}

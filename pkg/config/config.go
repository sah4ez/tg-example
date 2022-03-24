package config

import (
	"io"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const formatJSON = "json"

type ServiceConfig struct {
	LogLevel     string `envconfig:"LOGGER_LEVEL" default:"trace"`
	LogFormat    string `envconfig:"LOGGER_FORMAT" default:"console"`
	ReportCaller bool   `envconfig:"LOG_REPORT_CALLER" default:"false"`

	DSN string `envconfig:"DSN" default:"host=localhost port=5432 user=postgres dbname=postgres password=secretpass sslmode=disable client_encoding=UTF8"`

	Bind string `envconfig:"LISTEN" default:":9000"`

	BindMetrics string `envconfig:"LISTEN_METRICS" default:":9090"`
}

var service *ServiceConfig

func Service() ServiceConfig {

	if service != nil {
		return *service
	}
	service = &ServiceConfig{}
	if err := envconfig.Process("", service); err != nil {
		panic(err)
	}
	return *service
}

func (cfg ServiceConfig) Logger() (logger zerolog.Logger) {

	level := zerolog.InfoLevel
	if newLevel, err := zerolog.ParseLevel(cfg.LogLevel); err == nil {
		level = newLevel
	}
	var out io.Writer = os.Stdout
	if cfg.LogFormat != formatJSON {
		out = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	}
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	return zerolog.New(out).Level(level).With().Timestamp().Logger()
}

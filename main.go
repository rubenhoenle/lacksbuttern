package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	LACK = iota
	LACHS

	timeout = 5 * time.Second
)

func main() {
	var verbose bool
	var address, stringServerType string
	flag.BoolVar(&verbose, "v", false, "verbose output: sets the log level to debug")
	flag.StringVar(&address, "address", "0.0.0.0:80", "the address including port of the service")
	flag.StringVar(&stringServerType, "type", "", "the type of the server, either 'lack' or 'lachs'")
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, AddressCTX, address)
	if verbose {
		log.Ctx(ctx).Level(zerolog.DebugLevel)
	} else {
		log.Ctx(ctx).Level(zerolog.InfoLevel)
	}

	ctx = log.With().Logger().WithContext(ctx)

	switch stringServerType {
	case "lack":
		ctx = context.WithValue(ctx, TypeCTX, LACK)
	case "lachs":
		ctx = context.WithValue(ctx, TypeCTX, LACHS)
	case "":
		log.Ctx(ctx).Fatal().Msg("no server type was provided")
	}
	ctx = log.Ctx(ctx).With().Str("service", stringServerType).Logger().WithContext(ctx)

	server := NewServer(ctx)
	go server.Serve()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	cancel()
	os.Exit(0)
}

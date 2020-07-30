package main

import (
	"os"
	"os/signal"
	"pso2na-discord-go/internal/sonichi"
	"syscall"

	"github.com/rs/zerolog/log"
)

var (
	_sonichi sonichi.Controller
)

func main() {
	var err error
	// TODO: Source in config values via cmd line args
	_sonichi, err = sonichi.New(sonichi.Config{
		Token: "",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize sonichi :(")
	}

	if err := sonichi.BindHandlers(
		_sonichi,
		_sonichi.SayHello,
		_sonichi.Welcome,
	); err != nil {
		log.Fatal().Err(err).Msg("failed to bind event handlers")
	}

	if err := sonichi.Start(_sonichi); err != nil {
		log.Fatal().Err(err).Msg("failed to open websocket connection")
	}

	log.Info().Msg("Bot is now running, press CTRL+C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

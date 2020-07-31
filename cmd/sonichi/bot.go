package main

import (
	"errors"
	"fmt"
	"pso2na-discord-go/internal/sonichi"
	"strings"
)

func initSonichi() error {
	_botKey = strings.TrimSpace(_botKey)

	if _botKey == "" {
		return errors.New("botKey cannot be empty or whitespace")
	}

	var err error

	_sonichi, err = sonichi.New(sonichi.Config{
		Token: _botKey,
	})
	if err != nil {
		return fmt.Errorf("failed to create new discord bot, %w", err)
	}

	if err := sonichi.BindHandlers(
		_sonichi,
		_sonichi.SayHello,
	); err != nil {
		return fmt.Errorf("failed to bind event handlers, %w", err)
	}

	if err := sonichi.Start(_sonichi); err != nil {
		return fmt.Errorf("failed to open websocket connection, %w", err)
	}

	return nil
}

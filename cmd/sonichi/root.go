package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var _rootCmd = &cobra.Command{
	Use:   "sonichi [options]",
	Short: "Sonichi is a PSO2 NA Discord Bot!",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Debug().Msgf("no arguments found")
			if err := cmd.Help(); err != nil {
				log.Fatal().Err(err).Msg("failed to return help documentation")
			}
			os.Exit(0)
		}
	},
	TraverseChildren: true,
}

func execute() {
	if err := _rootCmd.Execute(); err != nil {
		os.Exit(0)
	}

	if err := initSonichi(); err != nil {
		log.Fatal().Err(err).Msg("failed to initialize sonichi :(")
	}

	log.Info().Msg("Bot is now running, press CTRL+C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func init() {
	_rootCmd.Flags().StringVar(&_botKey, "botKey", "", "The Discord Bot API Key")
	_rootCmd.MarkFlagRequired("botKey")
}

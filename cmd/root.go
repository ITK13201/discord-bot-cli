package cmd

import (
	"errors"
	"fmt"
	"github.com/ITK13201/discord-bot-cli/discord_bot_cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Channel struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Config struct {
	Token    string    `json:"token"`
	Channels []Channel `json:"channels"`
}

var (
	cfgFile string
	config  Config

	rootCmd = &cobra.Command{
		Use:   "discord-bot-cli",
		Short: "cli application for notifications to discord",
		Long:  "cli application for notifications to discord",
		Run: func(cmd *cobra.Command, args []string) {
			channelName, err := cmd.Flags().GetString("channel")
			channelID := ""
			if err == nil {
				for i := range config.Channels {
					if config.Channels[i].Name == channelName {
						channelID = config.Channels[i].ID
						break
					}
				}
				if channelID == "" {
					log.Fatal(errors.New(fmt.Sprintf("channel ID not found: %s", channelName)))
				}
			} else {
				log.Fatal(err)
			}
			level, err := cmd.Flags().GetString("level")
			if err != nil {
				log.Fatal(err)
			}
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				log.Fatal(err)
			}
			description, err := cmd.Flags().GetString("description")
			if err != nil {
				log.Fatal(err)
			}
			discord_bot_cli.Run(
				config.Token,
				channelID,
				level,
				title,
				description,
			)
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.discord-bot-cli.yaml)")
	rootCmd.PersistentFlags().StringP("channel", "c", "", "channel name defined in config file")
	rootCmd.PersistentFlags().StringP("level", "l", "info", "level (allowed: 'info', 'warn', 'error')")
	rootCmd.PersistentFlags().StringP("title", "t", "", "title")
	rootCmd.PersistentFlags().StringP("description", "d", "", "description")

	if err := rootCmd.MarkPersistentFlagRequired("channel"); err != nil {
		log.Fatal(err)
	}
	if err := rootCmd.MarkPersistentFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
	if err := rootCmd.MarkPersistentFlagRequired("description"); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".discord-bot-cli")
	}

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
}

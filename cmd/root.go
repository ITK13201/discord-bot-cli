package cmd

import (
	"errors"
	"github.com/ITK13201/discord-bot-cli/discord_bot_cli"
	"log"
	"os"
	"slices"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Token string `json:"token"`
}

var (
	cfgFile string
	config  Config
	levels  = []string{"info", "warn", "error"}

	rootCmd = &cobra.Command{
		Use:   "discord-bot-cli",
		Short: "cli application for notifications to discord",
		Long:  "cli application for notifications to discord",
		Run: func(cmd *cobra.Command, args []string) {
			channel, err := cmd.Flags().GetString("channel")
			if err != nil {
				log.Fatal(err)
			}
			level, err := cmd.Flags().GetString("level")
			if err == nil {
				if !slices.Contains(levels, level) {
					log.Fatal(errors.New("'level' flag must be one of these: 'info', 'warn', 'error'"))
				}
			} else {
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
				channel,
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
	rootCmd.PersistentFlags().StringP("channel", "c", "", "channel ID")
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

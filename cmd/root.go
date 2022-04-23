package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "command",
	Short: "command description",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	PreRun: func(_ *cobra.Command, _ []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {},
}

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

	rootCmd.PersistentFlags().String("log-level", "info", "trace, debug, info, error, warning, fatal, panic")
	_ = viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
}

func initConfig() {
	logLevel := viper.GetString("log-level")
	// validation
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Warningln("log-level", logLevel, "not exists. Setting up default 'info' level")
		level = log.InfoLevel
	}
	log.SetLevel(level)

	// display where the log was called
	if level >= log.DebugLevel {
		viper.SetDefault("debug", true)
		log.SetReportCaller(true)
	}
}

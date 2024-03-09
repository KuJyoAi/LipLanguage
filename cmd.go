package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"jcz-backend/internal/engine"
	"jcz-backend/model"
	"jcz-backend/router"
)

var rootCmd = &cobra.Command{
	Use: "jcz",
}

func init() {
	serverCmd.Flags().StringP("port", "p", "8080", "server port")
	serverCmd.Flags().StringP("config", "c", "", "config file")

	rootCmd.AddCommand(serverCmd)

	rootCmd.AddCommand(testCmd)

	rootCmd.AddCommand(migrateDBCmd)

	rootCmd.AddCommand(buildLearnCmd)
}

var migrateDBCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		// migrate db
		if err := engine.GetSqlCli().AutoMigrate(
			model.User{},
			model.UserLearnTime{},
		); err != nil {
			panic(err)
		}
	},
}

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		s := router.NewServer()
		port := cmd.Flag("port").Value.String()

		configPath := cmd.Flag("config").Value.String()
		if configPath != "" {
			viper.AddConfigPath(configPath)
		}

		logrus.SetReportCaller(true)

		if err := s.Run(":" + port); err != nil {
			panic(err)
		}
	},
}

var testCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		println("test")
	},
}

var buildLearnCmd = &cobra.Command{
	Use: "build_learn",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: build learn data
	},
}

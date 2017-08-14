package cmd

import (
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/gemnasium/logrus-airbrake-hook.v2"
	"github.com/enkhalifapro/pgen/db"
	"github.com/enkhalifapro/pgen/utilities"
	"github.com/enkhalifapro/pgen/server"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run server",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logrus.StandardLogger()
		if lvl, err := logrus.ParseLevel(viper.GetString("log.level")); err == nil {
			logger.Level = lvl
		}
		logger.Out = os.Stderr

		if viper.GetBool("airbrake.enable") {
			logger.Hooks.Add(
				airbrake.NewHook(
					int64(viper.GetInt("airbrake.id")),
					viper.GetString("airbrake.key"),
					viper.GetString("airbrake.env"),
				),
			)
		}

		// DB
		dbc, err := db.Dial(viper.GetString("db.uri"))
		if err != nil {
			logger.Error(err)
			return
		}
		defer dbc.Close()

		// Creates a gin router with default middleware:
		engine := gin.Default()
		if viper.GetBool("log.ginrus") {
			engine.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
		}

		server := &server.Server{}

		graph := &inject.Graph{}
		err = graph.Provide(
			&inject.Object{Value: dbc},
			&inject.Object{Value: &utilities.S3{}},
			&inject.Object{Value: &utilities.MandrillMailUtil{}},

			// Provide engine
			&inject.Object{Value: engine},
			&inject.Object{Value: server},
		)
		if err != nil {
			logger.Error(err)
			return
		}

		if err := graph.Populate(); err != nil {
			logger.Error(err)
			return
		}

		if err := server.Run(); err != nil {
			logger.Error(err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

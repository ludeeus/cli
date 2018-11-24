package cmd

import (
	"fmt"

	helper "github.com/home-assistant/hassio-cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// optionsCmd represents the options command
var homeassistantOptionsCmd = &cobra.Command{
	Use:     "options",
	Aliases: []string{"op"},
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("args", args).Debug("homeassistant options")

		section := "homeassistant"
		command := "options"
		base := viper.GetString("endpoint")

		options := make(map[string]interface{})

		for _, value := range []string{
			"image",
			"last_version",
			"password",
			"refresh_token",
		} {
			val, err := cmd.Flags().GetString(value)
			if val != "" && err == nil && cmd.Flags().Changed(value) {
				options[value] = val
			}
		}

		port, err := cmd.Flags().GetInt("port")
		if port != 0 && err == nil && cmd.Flags().Changed("port") {
			options["port"] = port
		}

		startupTime, err := cmd.Flags().GetInt("startup-time")
		if startupTime != 0 && err == nil && cmd.Flags().Changed("startup-time") {
			options["startup_time"] = startupTime
		}

		ssl, err := cmd.Flags().GetBool("ssl")
		if err == nil && cmd.Flags().Changed("ssl") {
			options["ssl"] = ssl
		}

		watchdg, err := cmd.Flags().GetBool("watchdg")
		if err == nil && cmd.Flags().Changed("watchdog") {
			options["watchdg"] = watchdg
		}

		request := helper.GetJSONRequest()
		if len(options) > 0 {
			log.WithField("options", options).Debug("Sending options")
			request.SetBody(options)
		}
		resp, err := request.Post(url)
		log.WithField("Request", resp.Request.RawRequest).Debug("Request")

		// returns 200 OK or 400
		if resp.StatusCode() != 200 && resp.StatusCode() != 400 {
			fmt.Println("Unexpected server response")
			fmt.Println(resp.String())
		} else {
			helper.ShowJSONResponse(resp)
		}

		return
	},
}

func init() {
	homeassistantOptionsCmd.Flags().String("image", "", "Optional image")
	homeassistantOptionsCmd.Flags().String("last_version", "", "Optional for custom image")
	homeassistantOptionsCmd.Flags().Int("port", 8123, "Port for access hassio")
	homeassistantOptionsCmd.Flags().Bool("ssl", false, "Use SSL")
	homeassistantOptionsCmd.Flags().String("password", "", "Api password")
	homeassistantOptionsCmd.Flags().String("refresh_token", "", "Refresh token")
	homeassistantOptionsCmd.Flags().Bool("watchdog", true, "Use watchdog")
	homeassistantOptionsCmd.Flags().Int("startup_time", 600, "startup time")
	homeassistantCmd.AddCommand(homeassistantOptionsCmd)
}
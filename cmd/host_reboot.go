package cmd

import (
	"fmt"

	helper "github.com/home-assistant/hassio-cli/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rebootCmd represents the reboot command
var hostRebootCmd = &cobra.Command{
	Use:     "reboot",
	Aliases: []string{"rb"},
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("args", args).Debug("host reboot")

		section := "host"
		command := "reboot"
		base := viper.GetString("endpoint")

		resp, err := helper.GenericJSONPost(base, section, command, nil)
		if err != nil {
			fmt.Println(err)
		} else {
			helper.ShowJSONResponse(resp)
		}
		return
	},
}

func init() {
	hostCmd.AddCommand(hostRebootCmd)
}
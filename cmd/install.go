/*
Copyright © 2023 Fleco Development shane@scaffoe.com
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/FlecoDevelopment/installer/ui"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Starts the installer for Fleco",
	Long:  `Installs Fleco onto your system from a git repo or a local folder with Docker or a Systemd service.`,
	Run: func(cmd *cobra.Command, args []string) {

		fromFlag, err := cmd.Flags().GetString("from")
		if err != nil {
			log.Fatalln(err)
		}

		if len(fromFlag) > 0 {

			fromDirInfo, err := os.Stat(fromFlag)
			if err != nil {
				log.Fatalln(err)
			}

			if !fromDirInfo.IsDir() {
				log.Fatalln(errors.New("--from must be a directory"))
			}

		}

		dockerInstallFlag, err := cmd.Flags().GetBool("docker-compose")
		if err != nil {
			log.Fatalln(err)
		}

		systemInstallFlag, err := cmd.Flags().GetBool("systemctl")
		if err != nil {
			log.Fatalln(err)
		}

		listItems := []ui.ListItem{
			{Label: "Docker Compose", Desc: "Install Fleco via docker-compose", Value: "docker"},
			{Label: "Systemd", Desc: "Install Fleco on the system using systemctl", Value: "systemd"},
		}

		var installMethod string

		if systemInstallFlag && dockerInstallFlag {

			installMethod = ui.NewList(listItems).(string)

		} else if !systemInstallFlag && !dockerInstallFlag {

			installMethod = ui.NewList(listItems).(string)

		} else {

			if systemInstallFlag {
				installMethod = "systemd"
			} else if dockerInstallFlag {
				installMethod = "docker"
			}

		}

		fmt.Println(installMethod)

	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")
	installCmd.PersistentFlags().String("from", "", "Specify the installation nto install from filesystem or from git repo.")
	installCmd.PersistentFlags().BoolP("docker-compose", "", false, "Install with docker-compose")
	installCmd.PersistentFlags().BoolP("systemctl", "", false, "Install with systemctl")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

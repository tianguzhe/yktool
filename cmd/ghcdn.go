package cmd

import (
	"errors"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var GhcdnCmd = &cobra.Command{
	Use:   "ghcdn",
	Short: "transfrom jsclide link",
	Args: func(cmd *cobra.Command, args []string) error {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			color.Red("%s", err.Error())
			return err
		}

		if !((strings.HasPrefix(url, "https://github.com") && strings.Contains(url, "blob")) ||
			strings.HasPrefix(url, "https://raw.githubusercontent.com")) {
			return errors.New("github address error")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		var cdn_url string

		url, _ := cmd.Flags().GetString("url")

		if strings.HasPrefix(url, "https://raw.githubusercontent.com") {

			var branchIndex = 5

			tmp := strings.Split(url, "/")

			cdn_url = strings.ReplaceAll(url, "https://raw.githubusercontent.com", "https://cdn.jsdelivr.net/gh")
			cdn_url = strings.ReplaceAll(cdn_url, "/"+tmp[branchIndex], "@"+tmp[branchIndex])

		} else {
			cdn_url = strings.ReplaceAll(url, "https://github.com", "https://cdn.jsdelivr.net/gh")
			cdn_url = strings.ReplaceAll(cdn_url, "/blob/", "@")
		}

		color.Green(cdn_url)
	},
}

func init() {
	GhcdnCmd.PersistentFlags().StringP("url", "u", "", "加速的github地址")
	rootCmd.AddCommand(GhcdnCmd)
}

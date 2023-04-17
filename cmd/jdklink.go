package cmd

import (
	"strings"
	"sync"
	"yktool/service"

	"github.com/bytedance/sonic"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type JdkPath struct {
	Binary Binary `json:"binary"`
}

type Installer struct {
	Link string `json:"link"`
	Name string `json:"name"`
	Size int    `json:"size"`
}

type Binary struct {
	Installer Installer `json:"installer"`
	UpdatedAt string    `json:"updated_at"`
}

var JdkLinkCmd = &cobra.Command{
	Use:   "jdklink",
	Short: "获取 openjdk 下载链接",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		softOs, _ := cmd.Flags().GetString("softOs")
		hardware, _ := cmd.Flags().GetString("hardware")
		version, _ := cmd.Flags().GetString("version")

		var ch = make(chan []byte, 3)
		var wg sync.WaitGroup

		var jdkVersion = []string{
			"8",
			"11",
			"17",
		}

		if version != "" {
			jdkVersion = []string{
				version,
			}
		}

		for i := 0; i < len(jdkVersion); i++ {
			wg.Add(1)

			go func(index int) {
				defer wg.Done()
				service.GetJdk(jdkVersion[index], ch, softOs)
			}(i)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		for ch := range ch {
			var jdkPath []JdkPath
			err := sonic.Unmarshal(ch, &jdkPath)

			if err != nil {
				color.Red("%v", err)
			}

			for _, value := range jdkPath {
				if strings.Contains(value.Binary.Installer.Name, hardware) {
					color.Blue("\n%s\n", value.Binary.Installer.Name)
					color.White(value.Binary.Installer.Link)
					color.Red(value.Binary.UpdatedAt)
				}
			}
		}
	},
}

func init() {
	JdkLinkCmd.PersistentFlags().StringP("softOs", "o", "mac", "系统版本 mac/windows")
	JdkLinkCmd.PersistentFlags().StringP("hardware", "x", "x64", "系硬件架构 x64/aarch64")
	JdkLinkCmd.PersistentFlags().StringP("version", "v", "", "需要的jdk版本 默认为8以后所有lts版本")
	rootCmd.AddCommand(JdkLinkCmd)
}

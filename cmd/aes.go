package cmd

import (
	"fmt"
	"os"
	"yktool/util"

	"github.com/spf13/cobra"
)

var AESCmd = &cobra.Command{
	Use:   "aes",
	Short: "文件加密",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("filePath")
		password, _ := cmd.Flags().GetString("password")
		model, _ := cmd.Flags().GetString("model")

		if filePath == "" || password == "" || model == "" {
			panic("请确认参数 -f 文件路径 -p 秘钥 -m 模式")
		}

		sourceFile, err := os.ReadFile(filePath)
		if err != nil {
			panic(err.Error())
		}

		for len(password) < 16 {
			password += "0"
		}

		key := []byte(password)

		if model == "encode" {
			encrypted, err := util.EncryptByAes(sourceFile, key)
			if err != nil {
				panic(err.Error())
			}

			outFile := fmt.Sprintf("%s.encode.txt", filePath)
			os.WriteFile(outFile, []byte(encrypted), 0644)

			fmt.Printf("加密文件为 filePath: %s, 秘钥为 %s", outFile, password)
		} else if model == "decode" {
			decrypted, err := util.DecryptByAes(string(sourceFile), key)
			if err != nil {
				panic(err.Error())
			}

			outFile := fmt.Sprintf("%s.decode.txt", filePath)
			os.WriteFile(outFile, []byte(decrypted), 0644)

			fmt.Printf("解密文件为 filePath: %s", outFile)
		} else {
			panic("未知模式")
		}
	},
}

func init() {
	AESCmd.PersistentFlags().StringP("filePath", "f", "", "加/解密文件路径")
	AESCmd.PersistentFlags().StringP("password", "p", "", "秘钥")
	AESCmd.PersistentFlags().StringP("model", "m", "", "encode/decode")
	rootCmd.AddCommand(AESCmd)
}

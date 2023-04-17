package cmd

import (
	"fmt"
	"net/url"
	"time"
	"yktool/util"

	"github.com/bytedance/sonic"
	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

type FinanceModel struct {
	Source string  `json:"source"`
	Target string  `json:"target"`
	Value  float64 `json:"value"`
	Time   int64   `json:"time"`
}

var FinanceCmd = &cobra.Command{
	Use:   "finance",
	Short: "实时汇率转换",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		quantity, _ := cmd.Flags().GetFloat64("quantity")
		source, _ := cmd.Flags().GetString("source")
		target, _ := cmd.Flags().GetString("target")
		filter, _ := cmd.Flags().GetInt("filter")

		baseUrl := "https://wise.com/rates/history+live"

		data := url.Values{}
		data.Set("source", source)
		data.Set("target", target)
		data.Set("length", "2")
		data.Set("resolution", "hourly")
		data.Set("unit", "day")

		result, err := util.GetWithParams(baseUrl, data)

		if err != nil {
			color.Red("api error: %s", err.Error())
		}

		var finances []FinanceModel

		err = sonic.Unmarshal(result, &finances)

		if err != nil {
			color.Red("json format error: %s", err.Error())
		}

		table := uitable.New()
		table.MaxColWidth = 100
		table.RightAlign(10)

		table.AddRow(
			"兑换数量", "持有货币", "当前汇率", "转换货币", "汇率更新时间",
		)

		if filter != 0 {
			finances = finances[len(finances)-filter:]
		}

		for _, finance := range finances {
			table.AddRow(
				color.MagentaString(fmt.Sprintf("%v", quantity)),
				color.RedString(finance.Source),
				color.GreenString(fmt.Sprintf("%v", finance.Value*quantity)),
				color.YellowString(finance.Target),
				time.Unix(finance.Time/1000, 0).Format("2006-01-02 15:04:05"),
			)
		}

		fmt.Println()
		fmt.Println(table)
	},
}

func init() {
	FinanceCmd.PersistentFlags().Float64P("quantity", "n", 1, "兑换金额")
	FinanceCmd.PersistentFlags().StringP("source", "s", "USD", "持有货币")
	FinanceCmd.PersistentFlags().StringP("target", "t", "CNY", "目标货币")
	FinanceCmd.PersistentFlags().IntP("filter", "f", 0, "目标货币")
	rootCmd.AddCommand(FinanceCmd)
}

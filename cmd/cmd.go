package cmd

import (
	"os"

	"github.com/antdate/antdate-service/pkg/log"
	"github.com/shuwenhe/shuwen-admin/configs"
	"github.com/shuwenhe/shuwen-admin/internal/admin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logger  = &logrus.Logger{}
	rootCmd = &cobra.Command{}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "configs/dev.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("debug", true, "开启debug")
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))

}

func Execute() error {
	rootCmd.AddCommand(CollyCommand)
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		admin.Dashboard()
		return nil
	}

	return rootCmd.Execute()

}

func initConfig() {
	// 配置初始化
	config.MustInit(os.Stdout, cfgFile)
	// 日志
	logger = log.New()
}

package spider

import "github.com/spf13/cobra"

var confPath string

func SpiderCommand(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&confPath, "config", "c", "./config.conf", "config file (default is ./config.conf)")
	NewLive(cmd)
	NewJuxing(cmd)
	NewTest(cmd)
}

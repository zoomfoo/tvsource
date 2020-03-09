package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zoomfoo/iplaytv/spider"
)

var cmdServer = &cobra.Command{
	Use:   "iplaytv",
	Short: "iplaytv",
	Long:  `iplaytv`,
	Run: func(_ *cobra.Command, _ []string) {
		logrus.Info("正在执行")
		spider.Crawler()
	},
}

func main() {
	cmdRoot := &cobra.Command{Use: "iplaytv", Version: "0.0.1"}
	cmdRoot.AddCommand(cmdServer)
	cmdRoot.Execute()
}

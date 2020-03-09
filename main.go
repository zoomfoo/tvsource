package main

import (
	"github.com/spf13/cobra"
	"github.com/zoomfoo/tvsource/spider"
)

func main() {
	cmdRoot := &cobra.Command{Use: "iplaytv", Version: "0.0.1"}
	cmdRoot.AddCommand(spider.CmdServer)
	cmdRoot.Execute()
}

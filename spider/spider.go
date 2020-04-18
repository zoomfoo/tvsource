package spider

import "github.com/spf13/cobra"

func SpiderCommand(cmd *cobra.Command) {
	NewDouyu(cmd)
	NewJuxing(cmd)
	NewTest(cmd)
}

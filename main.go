package main

// import (
// 	"github.com/spf13/cobra"
// 	"github.com/zoomfoo/tvsource/spider"
// )

// func main() {
// 	root := &cobra.Command{Use: "iplaytv", Version: "0.0.1"}
// 	spider.SpiderCommand(root)
// 	root.Execute()
// }

// Command visible is a chromedp example demonstrating how to wait until an
// element is visible.

// Command visible is a chromedp example demonstrating how to wait until an
// element is visible.

import (
	"github.com/spf13/cobra"
	"github.com/zoomfoo/tvsource/spider"
)

func main() {
	root := &cobra.Command{Use: "tvsource", Version: "0.0.1"}
	spider.SpiderCommand(root)
	root.Execute()
}

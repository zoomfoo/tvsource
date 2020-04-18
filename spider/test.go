package spider

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

func NewTest(cmd *cobra.Command) {
	testCmd := &cobra.Command{
		Use:   "test",
		Short: "test",
		Run: func(_ *cobra.Command, ids []string) {
			t := Test{
				liveIDs: ids,
			}
			t.findLivePage()
		},
	}
	cmd.AddCommand(testCmd)
}

type Test struct{ liveIDs []string }

func (t *Test) findLivePage() {
	ctx_, cancel_ := chromedp.NewExecAllocator(context.Background())
	defer cancel_()
	ctx, cancel := chromedp.NewContext(ctx_)
	defer cancel()

	c := chromedp.FromContext(ctx)

	chromedp.ListenTarget(cdp.WithExecutor(ctx, c.Target), func(event interface{}) {
		switch ev := event.(type) {
		case *network.EventResponseReceived:
			_ = ev
			// println(ev.Response.URL, "   ", ev.Response.MimeType)
		}
	})
	chromedp.Run(ctx, tasks())

	for {
		res := ""
		chromedp.Run(ctx, chromedp.Tasks{
			chromedp.OuterHTML("#publicBox", &res),
		})
		fmt.Println(res)
		time.Sleep(time.Second * 2)
	}
}

func tasks() chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		// chromedp.Emulate(device.IPhoneX),
		chromedp.Navigate("https://jx.kuwo.cn/690511?entrance=107"),
		// chromedp.Sleep(30 * time.Second),
		chromedp.Focus("#inputChatMessage"),
		chromedp.SetValue("#inputChatMessage", "你好"),
		chromedp.Click("#msgbtn"),
	}
}

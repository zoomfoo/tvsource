package spider

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
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
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	chromedp.ListenTarget(c, func(event interface{}) {
		switch ev := event.(type) {
		case *network.EventResponseReceived:
			_ = ev
			// println(ev.Response.URL, "   ", ev.Response.MimeType)
		}
	})

	ctx2, _ := chromedp.NewContext(ctx)

	// ensure the second tab is created
	if err := chromedp.Run(ctx2, chromedp.Tasks{chromedp.Navigate("https://ww.baidu.com")}); err != nil {
		panic(err)
	}

	chromedp.Run(ctx, tasks())

	for {
		data := []byte("")
		title := ""
		chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Title(&title),
			chromedp.CaptureScreenshot(&data),
		})
		fmt.Println(title)
		if len(data) > 0 {
			ioutil.WriteFile("wy.png", data, os.ModePerm)
		}
		time.Sleep(time.Second * 2)
	}
}

func tasks() chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		// chromedp.Emulate(device.IPhoneX),

		chromedp.Navigate("https://www.douyu.com/5377024"),
		// chromedp.Sleep(30 * time.Second),
		// chromedp.Focus("#inputChatMessage"),
		// chromedp.SetValue("#inputChatMessage", "你好"),
		// chromedp.Click("#msgbtn"),
	}
}

func SetCookie(name, value, domain, path string, httpOnly, secure bool) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
		success, err := network.SetCookie(name, value).
			WithExpires(&expr).
			WithDomain(domain).
			WithPath(path).
			WithHTTPOnly(httpOnly).
			WithSecure(secure).
			Do(ctx)
		if err != nil {
			return err
		}
		if !success {
			return fmt.Errorf("could not set cookie %s", name)
		}
		return nil
	})
}

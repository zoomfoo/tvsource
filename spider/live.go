package spider

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zoomfoo/iplaytv/httplib"
	"github.com/zoomfoo/tvsource/config"
)

const (
	DouyuHost = "https://m.douyu.com/"
)

func NewLive(cmd *cobra.Command) {
	liveCmd := &cobra.Command{
		Use:   "live",
		Short: "直播源获取",
		Long:  `执行一次主播 ID 直播源`,
		Run: func(_ *cobra.Command, _ []string) {
			conf := config.NewConfigure(confPath)
			if len(conf.DouyuLive) == 0 {
				logrus.Info("没有发现主播房间 id")
				os.Exit(1)
			}
			ids := make([]string, len(conf.DouyuLive))
			for _, v := range conf.DouyuLive {
				ids = append(ids, v)
			}
			dy := LiveBox{
				liveIDs: ids,
			}
			dy.run()
		},
	}
	cmd.AddCommand(liveCmd)
}

type LiveBox struct {
	liveIDs []string
}

func (lb *LiveBox) run() {
	lb.findLiveMu38()
}

func (lb *LiveBox) findLiveMu38() {
	for _, id := range lb.liveIDs {
		lb.findLivePage(id)
	}
}

func (lb *LiveBox) findLivePage(id string) {
	// ctx := context.Background()
	// options := []chromedp.ExecAllocatorOption{
	// 	chromedp.Flag("headless", true),
	// 	chromedp.Flag("hide-scrollbars", false),
	// 	chromedp.Flag("mute-audio", false),
	// }
	// options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	// c, cc := chromedp.NewExecAllocator(ctx, options...)

	// 使用前需安装 docker run -d -p 9222:9222 --rm --name headless-shell chromedp/headless-shell
	addr := "http://localhost:9222/json"
	simpleJson, err := httplib.Get(addr).ToSimpleJson()
	if err != nil {
		logrus.Errorf("get addr error. [err='%v']", err)
		os.Exit(1)
	}
	allocatorCtx, allocatorCancel := chromedp.NewRemoteAllocator(context.Background(), simpleJson.GetIndex(0).Get("webSocketDebuggerUrl").MustString())
	defer allocatorCancel()
	// create context
	ctx, cancel := chromedp.NewContext(allocatorCtx)
	defer cancel()

	lb.listenM3u8File(ctx)
	html := ""
	if err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.Emulate(device.IPhoneX),
		chromedp.Navigate(DouyuHost+id),
		chromedp.WaitReady("root", chromedp.ByID),
		chromedp.OuterHTML("html", &html),
	); err != nil {
		fmt.Println(err)
	}

	lb.printM3u8(goquery.NewDocumentFromReader(bytes.NewReader([]byte(html))))
}

func (lb *LiveBox) listenM3u8File(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(event interface{}) {
		switch ev := event.(type) {
		case *network.EventResponseReceived:
			_ = ev.Response.URL
		}
	})
}

func (lb *LiveBox) printM3u8(doc *goquery.Document, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	src, _ := doc.Find("#html5player-video").Attr("src")
	style, _ := doc.Find("#root > div > div.l-video > div > div.room-video-play > div > div > div").Attr("style")
	logo := getLogo(style)
	if logo == "" {
		return
	}
	fmt.Printf(`#EXTINF:-1 tvg-logo="%v" , %v`, logo, doc.Find("title").Text())
	fmt.Println("")
	fmt.Println(src)
}

func getLogo(str string) string {
	re, _ := regexp.Compile(`url\(.*"\)`)
	raw := re.FindString(str)
	for _, v := range []string{`url`, `"`, `(`, `)`, `;`} {
		raw = strings.Replace(raw, v, "", -1)
	}
	return strings.TrimSpace(raw)
}

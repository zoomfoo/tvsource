package spider

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zoomfoo/iplaytv/httplib"
)

const (
	DouyuHost = "https://m.douyu.com/"
)

func NewDouyu(cmd *cobra.Command) {
	douyuCmd := &cobra.Command{
		Use:   "douyu",
		Short: "douyu 直播源获取",
		Long:  `执行一次 斗鱼 主播 ID 直播源`,
		Run: func(_ *cobra.Command, ids []string) {
			dy := Douyu{
				liveIDs: ids,
			}
			dy.run()
		},
	}
	cmd.AddCommand(douyuCmd)
}

type Douyu struct {
	liveIDs []string
}

func (dy *Douyu) run() {
	dy.findLiveMu38()
}

func (dy *Douyu) findLiveMu38() {
	for _, id := range dy.liveIDs {
		dy.findLivePage(id)
	}
}

func (dy *Douyu) findLivePage(id string) {
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
		return
	}
	c, cc := chromedp.NewRemoteAllocator(context.Background(), simpleJson.GetIndex(0).Get("webSocketDebuggerUrl").MustString())
	defer cc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	dy.listenM3u8File(ctx)
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

	dy.printM3u8(goquery.NewDocumentFromReader(bytes.NewReader([]byte(html))))
}

func (dy *Douyu) listenM3u8File(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(event interface{}) {
		switch ev := event.(type) {
		case *network.EventResponseReceived:
			_ = ev.Response.URL
		}
	})
}

func (dy *Douyu) printM3u8(doc *goquery.Document, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	src, _ := doc.Find("#html5player-video").Attr("src")
	style, _ := doc.Find("#root > div > div.l-video > div > div.room-video-play > div > div > div").Attr("style")
	fmt.Printf(`#EXTINF:-1 tvg-logo="%v" , %v`, getLogo(style), doc.Find("title").Text())
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

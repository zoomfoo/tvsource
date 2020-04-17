package spider

import (
	"bytes"
	"context"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	JuxingHost = "https://jx.kuwo.cn/"
)

func NewJuxing(cmd *cobra.Command) {
	juxingCmd := &cobra.Command{
		Use:   "juxing",
		Short: "juxing 直播源获取",
		Long:  `执行一次聚星主播 ID 直播源`,
		Run: func(_ *cobra.Command, ids []string) {
			jx := Juxing{
				liveIDs: ids,
			}
			jx.run()
		},
	}
	cmd.AddCommand(juxingCmd)
}

type Juxing struct {
	liveIDs []string
}

func (jx *Juxing) run() {
	jx.findLiveMu38()
}

func (jx *Juxing) findLiveMu38() {
	for _, id := range jx.liveIDs {
		doc, err := jx.findLivePage(id)
		if err != nil {
			logrus.Error("findLiveMu38 err ", err)
			continue
		}
		jx.printM3u8(doc)
	}
}

func (jx *Juxing) findLivePage(id string) (*goquery.Document, error) {
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1`),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	var html = ""
	err := chromedp.Run(ctx,
		chromedp.Navigate(JuxingHost+id),
		chromedp.OuterHTML(`html`, &html),
	)
	if err != nil {
		return nil, err
	}
	// fmt.Println(html)
	return goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
}

func (jx *Juxing) printM3u8(doc *goquery.Document) {
	src, _ := doc.Find("#myVideo").Attr("src")
	fmt.Printf(`#EXTINF:-1 tvg-logo="%v" , %v\n`, "", doc.Find(".zhubo_name").Text())
	fmt.Println(src)
}

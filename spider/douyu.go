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
		doc, err := dy.findLivePage(id)
		if err != nil {
			logrus.Error("findLiveMu38 err ", err)
			continue
		}
		dy.printM3u8(doc)
	}
}

func (dy *Douyu) findLivePage(id string) (*goquery.Document, error) {
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
		chromedp.Navigate(DouyuHost+id),
		chromedp.WaitVisible(`#root`),
		chromedp.OuterHTML(`html`, &html),
	)
	if err != nil {
		return nil, err
	}
	// fmt.Println(html)
	return goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
}

func (dy *Douyu) printM3u8(doc *goquery.Document) {
	src, _ := doc.Find("#html5player-video").Attr("src")
	fmt.Printf("#EXTINF:-1 , %v\n", doc.Find("title").Text())
	fmt.Println(src)
}

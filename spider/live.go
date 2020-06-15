package spider

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zoomfoo/iplaytv/httplib"
	"github.com/zoomfoo/tvsource/config"
	"github.com/zoomfoo/tvsource/utils"
)

const (
	DouyuHost = "https://m.douyu.com/"
)

func NewLive(cmd *cobra.Command) {
	local := false
	liveCmd := &cobra.Command{
		Use:   "live",
		Short: "直播源获取",
		Long:  `执行一次主播 ID 直播源`,
		Run: func(_ *cobra.Command, ps []string) {
			conf := config.NewConfigure(confPath)
			if len(conf.DouyuLive) == 0 {
				logrus.Info("没有发现主播房间 id")
				os.Exit(1)
			}
			ids := make([]string, len(conf.DouyuLive))
			for _, v := range conf.DouyuLive {
				ids = append(ids, v)
			}
			cnt := 1
			if len(ps) > 0 {
				tmp := utils.StrTo(ps[0]).MustInt()
				if tmp != 0 {
					cnt = tmp
				}
			}
			if err := InitLiveBox(ids, cnt, local); err != nil {
				logrus.Error(err)
			}
		},
	}
	liveCmd.Flags().BoolVarP(&local, "local", "", false, "whether to use a local browser")
	cmd.AddCommand(liveCmd)
}

func InitLiveBox(ids []string, cnt int, local bool) error {
	box := &LiveBox{
		cnt:      cnt,
		urlsChan: make(chan string, cnt),
	}
	if local {
		if err := box.chromedpExecAllocator(); err != nil {
			return err
		}
	} else {
		if err := box.chromedpRemoteAllocatorRemote(); err != nil {
			return err
		}
	}

	go func() {
		for _, id := range ids {
			box.urlsChan <- DouyuHost + id
		}
	}()
	box.init()
	for {
		time.Sleep(time.Second)
		if len(box.urlsChan) == 0 {
			box.CloseUrlsChan()
			if box.done == int32(box.cnt) {
				box.Close()
				break
			}
		}
	}
	return nil
}

type LiveBox struct {
	cnt        int
	done       int32
	urlsChan   chan string
	urlsOnce   sync.Once
	outputLock sync.Mutex
	baseCtx    context.Context
	baseCancel func()
	cancels    []func()
}

func (lb *LiveBox) chromedpExecAllocator() error {
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	lb.baseCtx, lb.baseCancel = chromedp.NewExecAllocator(ctx, options...)
	return nil
}

func (lb *LiveBox) chromedpRemoteAllocatorRemote() error {
	addr := "http://localhost:9222/json"
	simpleJson, err := httplib.Get(addr).ToSimpleJson()
	if err != nil {
		return err
	}
	lb.baseCtx, lb.baseCancel = chromedp.NewRemoteAllocator(context.Background(), simpleJson.GetIndex(0).Get("webSocketDebuggerUrl").MustString())
	return nil
}

func (lb *LiveBox) init() {
	for i := 0; i < lb.cnt; i++ {
		ctx, cancel := chromedp.NewContext(lb.baseCtx)
		lb.cancels = append(lb.cancels, cancel)
		go lb.group(ctx, lb.urlsChan)
	}
}

func (lb *LiveBox) Close() {
	for _, cancel := range lb.cancels {
		cancel()
	}
}

func (lb *LiveBox) CloseUrlsChan() {
	lb.urlsOnce.Do(func() {
		close(lb.urlsChan)
	})
}

func (lb *LiveBox) group(ctx context.Context, urlsChan <-chan string) {
	defer atomic.AddInt32(&lb.done, 1)

	lb.listenM3u8File(ctx)
	for url := range urlsChan {
		html := ""
		if err := chromedp.Run(ctx,
			network.Enable(),
			chromedp.Emulate(device.IPhoneX),
			chromedp.Navigate(url),
			chromedp.WaitReady("root", chromedp.ByID),
			chromedp.OuterHTML("html", &html),
		); err != nil {
			fmt.Println(err)
			return
		}
		lb.printM3u8(goquery.NewDocumentFromReader(bytes.NewReader([]byte(html))))
	}
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
	lb.outputLock.Lock()
	fmt.Printf(`#EXTINF:-1 tvg-logo="%v" , %v`, logo, doc.Find("title").Text())
	fmt.Println("")
	fmt.Println(src)
	lb.outputLock.Unlock()
}

func getLogo(str string) string {
	re, _ := regexp.Compile(`url\(.*"\)`)
	raw := re.FindString(str)
	for _, v := range []string{`url`, `"`, `(`, `)`, `;`} {
		raw = strings.Replace(raw, v, "", -1)
	}
	return strings.TrimSpace(raw)
}

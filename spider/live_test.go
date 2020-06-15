package spider

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

func TestLive(t *testing.T) {
	lb := &LiveBox{}
	id := "96291"

	c, cc := lb.chromedpExecAllocator()
	defer cc()

	ctx1, cancel1 := chromedp.NewContext(c)
	defer cancel1()
	lb.listenM3u8File(ctx1)

	ctx2, cancel2 := chromedp.NewContext(c)
	defer cancel2()
	lb.listenM3u8File(ctx2)

	ctx3, cancel3 := chromedp.NewContext(c)
	defer cancel3()
	lb.listenM3u8File(ctx3)

	go func() {
		html1 := ""
		if err := chromedp.Run(ctx1,
			network.Enable(),
			chromedp.Emulate(device.IPhoneX),
			chromedp.Navigate(DouyuHost+"7918091"),
			chromedp.WaitReady("root", chromedp.ByID),
			chromedp.OuterHTML("html", &html1),
		); err != nil {
			fmt.Println(err)
			return
		}
		lb.printM3u8(goquery.NewDocumentFromReader(bytes.NewReader([]byte(html1))))
	}()

	go func() {
		html := ""
		if err := chromedp.Run(ctx2,
			network.Enable(),
			chromedp.Emulate(device.IPhoneX),
			chromedp.Navigate(DouyuHost+id),
			chromedp.WaitReady("root", chromedp.ByID),
			chromedp.OuterHTML("html", &html),
		); err != nil {
			fmt.Println(err)
			return
		}
		lb.printM3u8(goquery.NewDocumentFromReader(bytes.NewReader([]byte(html))))
	}()

	go func() {
		html3 := ""
		if err := chromedp.Run(ctx3,
			network.Enable(),
			chromedp.Emulate(device.IPhoneX),
			chromedp.Navigate(DouyuHost+"2222"),
			chromedp.WaitReady("root", chromedp.ByID),
			chromedp.OuterHTML("html", &html3),
		); err != nil {
			fmt.Println(err)
			return
		}
		lb.printM3u8(goquery.NewDocumentFromReader(bytes.NewReader([]byte(html3))))
	}()

	time.Sleep(time.Second * 200)
}

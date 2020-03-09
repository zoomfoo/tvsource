package spider

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"github.com/zoomfoo/iplaytv/httplib"
	"github.com/zoomfoo/iplaytv/utils"
)

const (
	host = "http://m.66zhibo.net"
)

var headers = map[string]string{
	"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
}

func Crawler() {
	wg := sync.WaitGroup{}
	keys := []string{"1", "2", "3", "4"}
	wg.Add(len(keys))
	for _, v := range keys {
		go func(key string) {
			defer func() { wg.Done() }()

			doc, err := getPage(fmt.Sprintf("%v/%v/", host, key))
			if err != nil {
				logrus.Error("NewDocumentFromReader err: ", err)
				return
			}

			doc.Find("body > div.wrap > div.list-box.J-medal > ul:nth-child(3) > li").Each(func(i int, s *goquery.Selection) {
				url, ok := s.Find("a").Attr("href")
				if !ok {
					return
				}
				title, _ := s.Find("a").Attr("title")
				data, _ := utils.GbkToUtf8([]byte(title))
				parseSource(host+url, string(data))
			})
		}(v)
	}
	wg.Wait()
}

func getPage(url string) (*goquery.Document, error) {
	ret := httplib.Get(url)
	for k, v := range headers {
		ret.Header(k, v)
	}
	data, err := ret.Bytes()
	if err != nil {
		logrus.Error("crawler err: ", err)
		return nil, err
	}
	return goquery.NewDocumentFromReader(bytes.NewReader(data))
}

func parseSource(url, title string) {
	doc, err := getPage(url)
	if err != nil {
		logrus.Error("NewDocumentFromReader err: ", err)
		return
	}

	html, err := doc.Find("body > section > div.play-bx > div script").Html()
	if err != nil {
		logrus.Error("Find err: ", err)
		return
	}

	html = strings.Replace(html, "&#39;", "", -1)

	sourIdRegex, err := regexp.Compile(`sourid=\d+`)
	if err != nil {
		logrus.Error("Find err: ", err)
		return
	}
	gidRegex, err := regexp.Compile(`gid=\d+`)
	if err != nil {
		logrus.Error("Find err: ", err)
		return
	}
	vRegex, err := regexp.Compile(`v=\w+`)
	if err != nil {
		logrus.Error("Find err: ", err)
		return
	}

	sourid := strings.Replace(sourIdRegex.FindString(html), "sourid=", "", -1)
	gid := strings.Replace(gidRegex.FindString(html), "gid=", "", -1)
	v := strings.Replace(vRegex.FindString(html), "v=", "", -1)

	parseVideo(sourid, gid, v, title)
}

func parseVideo(sourceId, gid, v, title string) {
	url := fmt.Sprintf("%v/e/extend/tv.php?id=%v&gid=%v&v=%v", host, sourceId, gid, v)
	doc, err := getPage(url)
	if err != nil {
		logrus.Error("parseVideo err ", err)
		return
	}

	html, err := doc.Html()
	if err != nil {
		logrus.Error("doc.html err ", err)
		return
	}

	re, err := regexp.Compile(`\$(.*)\$`)
	if err != nil {
		logrus.Error("regexp.Compile err ", err)
		return
	}
	data := strings.Replace(re.FindString(html), "$", "", -1)
	fmt.Println("#EXTINF:-1 ,", title)
	fmt.Println(data)
}

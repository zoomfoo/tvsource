package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GetMD5HashString func
func GetMD5HashString(str string) string {
	return GetMD5HashBytes([]byte(str))
}

// GetMD5HashBytes func
func GetMD5HashBytes(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

// GetRandomString func
func GetRandomString(n int, alphabets ...byte) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes)
}

// Errors func
func Errors(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// WaitTimeout fn
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(timeout):
		return true
	}
}

// DeferError fn
func DeferError(errorfn func(string, interface{}), dones ...func()) {
	if err := recover(); err != nil {
		var buf [2 << 10]byte
		errorfn(string(buf[:runtime.Stack(buf[:], false)]), err)
	}
	for _, done := range dones {
		done()
	}
}

// ParseIPAndPort func
func ParseIPAndPort(addr string) (ip string, port string) {
	strs := strings.Split(addr, ":")
	if len(strs) < 2 {
		return
	}
	ip, port = strs[0], strs[1]
	return
}

// GbkToUtf8 func
func GbkToUtf8(bs []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// GbkToUtf8String func
func GbkToUtf8String(s string) (string, error) {
	data, err := GbkToUtf8([]byte(s))
	return string(data), err
}

// Append func
func Append(slice []string, data ...string) []string {
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]string, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

const (
	WeightDel uint8 = 1
)

type item struct {
	value  interface{}
	weight int
	status uint8
}

type Weight struct {
	items         []*item
	rm            bool
	currItemValue interface{}
}

func (w *Weight) Add(k string, v interface{}, weight int) *Weight {
	if w.items == nil {
		w.items = []*item{}
	}
	w.items = append(w.items, &item{value: v, weight: weight})
	return w
}

func (w *Weight) isItemRM(i *item) bool {
	return w.rm && i.status == WeightDel
}

func (w *Weight) RandomValue() interface{} {
	total := 0
	for _, item := range w.items {
		if w.isItemRM(item) {
			continue
		}
		total += item.weight
	}
	if total == 0 {
		return nil
	}
	rd := rand.Intn(total)
	currsum := 0
	for _, item := range w.items {
		if w.isItemRM(item) {
			continue
		}
		currsum += item.weight
		if rd <= currsum {
			if w.rm {
				item.status = WeightDel
			}
			return item.value
		}
	}
	return nil
}

func (w *Weight) NextRandom() bool {
	w.rm = true
	w.currItemValue = w.RandomValue()
	return w.currItemValue != nil
}

func (w *Weight) Value() interface{} {
	return w.currItemValue
}

func CheckIP(addr string) error {
	conn, err := net.DialTimeout("tcp", addr, time.Second*5)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	return err
}

const (
	formatStart = iota
	formatAllAnd
	formatAllOr
	formatGroupStart
	formatGroupEnd
	formatGroupEndAnd
)

// TopicParse func
func TopicParse(str string) ([][]string, error) {
	state := formatStart
	off := 0
	base := [][]string{}
	group := []string{}
	last := len(str) - 1
loop:
	for i, s := range str {
		switch state {
		case formatStart:
			if unicode.IsLetter(s) {
				if i == last {
					group = Append(group, str[off:])
					base = append(base, group)
				}
				continue loop
			}
			switch s {
			case '|':
				state = formatAllOr
				group = append(group, str[off:i])
				off = i + 1
				continue loop
			case ',':
				state = formatAllAnd
				base = append(base, []string{str[off:i]})
				off = i + 1
				continue loop
			}
			if s == '(' {
				state = formatGroupStart
				off = i + 1
			}

		case formatAllAnd:
			if s == ',' {
				base = append(base, []string{str[off:i]})
				off = i + 1
				continue loop
			}

			if s == '(' {
				group = []string{}
				state = formatGroupStart
				off = i + 1
				continue
			}

			if i == last {
				group = append(group, str[off:])
				break
			}

		case formatAllOr:
			switch s {
			case '(':
				return nil, fmt.Errorf("state formatAllOr: %v %v", str, "Unsupported format")
			}

			if s == '|' {
				group = append(group, str[off:i])
				off = i + 1
				continue loop
			}

			if s == ')' {
				state = formatGroupEnd
				group = append(group, str[off:i])
				base = append(base, group)
				off = i + 1
				continue loop
			}

			if i == last {
				group = append(group, str[off:])
				break
			}

		case formatGroupStart:
			if unicode.IsLetter(s) {
				continue loop
			}

			if s == ')' {
				state = formatGroupEnd
				group = append(group, str[off:i])
				base = append(base, group)
				off = i + 1
				continue loop
			}

			switch s {
			case ',':
				return nil, fmt.Errorf("%v %v", str, "Commas in groups are not supported for the time being")
			case '|':
				state = formatAllOr
				group = Append(group, str[off:i])
				off = i + 1
				continue loop
			}

		case formatGroupEnd:
			if unicode.IsSpace(s) {
				continue loop
			}
			if s == ',' {
				state = formatGroupEndAnd
				off = i + 1
				continue loop
			}
			return nil, fmt.Errorf("state formatGroupEnd: %v %v", str, "Unsupported format")

		case formatGroupEndAnd:
			if unicode.IsLetter(s) {
				state = formatStart
				group = []string{}
				if i == last {
					group = Append(group, str[off:])
					base = append(base, group)
				}
				continue loop
			}
			if s == '(' {
				group = []string{}
				state = formatGroupStart
				off = i + 1
				continue loop
			}
		}

	}
	switch state {
	case formatAllOr:
		base = append(base, group)
	case formatAllAnd:
		for i := range group {
			base = append(base, []string{group[i]})
		}
	}
	return base, nil
}

var loc, _ = time.LoadLocation("Asia/Shanghai")

func TimeParseInShanghai(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}

func ShanghaiNowTime() time.Time { return time.Now().In(loc) }

package spider

import (
	"testing"
)

func TestGetLogo(t *testing.T) {
	t.Log(getLogo(`background-image: url("https://rpic.douyucdn.cn/asrpic/200418/252140_0101.png/dy1");`))
}

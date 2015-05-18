package crawler

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
)

func TestConver(t *testing.T) {
	//花间一壶酒，独酌无相亲。
	var src = "\xbb\xa8\xbc\xe4\xd2\xbb\xba\xf8\xbe\xc6\xa3\xac\xb6\xc0\xd7\xc3\xce\xde\xcf\xe0\xc7\xd7\xa1\xa3"
	decode := simplifiedchinese.GBK.NewDecoder()
	s := Conver(src, decode)
	fmt.Printf("conver:  %s\n", s)
}

func TestCrawler(t *testing.T) {
	//var url = "http://finance.sina.com.cn/china/hgjj/20150427/023422048971.shtml"
	//var container = ".blkContainerSblk"
	//text := Crawler(url, container, "gb2312")
	var url = "http://www.woshipm.com/pmd/151979.html"
	var container = ".con_txt.clx"
	text := Crawler(url, container, "utf-8")
	fmt.Printf("conver:  %s\n", text)
}

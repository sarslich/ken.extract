package cacul

import (
	"fmt"
	"ken.extract/crawler"
	"testing"
)

func TestTextExtract(t *testing.T) {
	//	var url = "http://finance.sina.com.cn/china/hgjj/20150427/023422048971.shtml"
	//	var container = ".blkContainerSblk"
	//	text := crawler.Crawler(url, container, "gb2312")

	var url = "http://www.woshipm.com/pmd/151979.html"
	var container = ".con_txt.clx"
	text := crawler.Crawler(url, container, "utf-8")
	fmt.Printf("source Text:  %s\n", text)
	endText := TextExtract(text, false, 85)
	fmt.Printf("TextExtract:  %s\n", endText)
}

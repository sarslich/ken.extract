package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"strings"
)

func Conver(src string, decoder transform.Transformer) string {
	r := strings.NewReader(src)
	rInUtf8 := transform.NewReader(r, decoder)
	g, err := ioutil.ReadAll(rInUtf8)
	if err != nil {
		log.Fatal(err)
	}
	output := string(g)
	return output
}

func Crawler(url string, container string, encode string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	sel := doc.Find(container)
	//filter script,style tag
	sel.Find("script").Remove()
	sel.Find("style").Remove()

	text := sel.Text()

	//encode conver
	var decode transform.Transformer
	switch encode {
	case "gbk", "GBK", "gb2312", "GB2312":
		decode = simplifiedchinese.GBK.NewDecoder()
		text = Conver(text, decode)
	}

	return text

}

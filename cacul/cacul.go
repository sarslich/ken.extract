package cacul

import (
	//	"fmt"
	//	"log"
	"strings"
)

const blocksWidth = 3 //行块厚度，通常<5,表示行块由几行构成；

/* 当待抽取的网页正文中遇到成块的新闻标题未剔除时，只要增大此阈值(threshold)即可。*/
/* 阈值增大，准确率提升，召回率下降；值变小，噪声会大，但可以保证抽到只有一句话的正文 */
func TextExtract(src string, flag bool, threshold int) string {

	lines := strings.Split(src, "\n")

	//	empty := 0 //空行的数量

	lenOfLines := len(lines)

	indexDistribution := make([]int, lenOfLines)

	//数据预处理，消除空格等字符
	for i := 0; i < lenOfLines; i++ {
		lines[i] = replaceBlank(lines[i])
	}

	for i := 0; i < lenOfLines-blocksWidth; i++ {
		//		if len(lines[i]) == 0 {
		//			empty++
		//		}
		wordsNum := 0
		//行块密度计算；
		for j := i; j < i+blocksWidth && j < lenOfLines; j++ {
			wordsNum += len(lines[j])
		}
		indexDistribution[i] = wordsNum
	}

	text := ""
	start, end := -1, -1
	boolstart, boolend := false, false
	firstMatch := true //前面的标题块往往比较小，应该减小与它匹配的阈值

	for i := 0; i < len(indexDistribution)-1; i++ {

		if firstMatch && !boolstart {
			if indexDistribution[i] > (threshold/2) && !boolstart {
				if indexDistribution[i+1] != 0 || indexDistribution[i+2] != 0 {
					firstMatch = false
					boolstart = true
					start = i
					continue
				}
			}
		}

		if indexDistribution[i] > threshold && !boolstart {
			if indexDistribution[i+1] != 0 || indexDistribution[i+2] != 0 || indexDistribution[i+3] != 0 {
				boolstart = true
				start = i
				continue
			}
		}

		if boolstart {
			if indexDistribution[i] == 0 || indexDistribution[i+1] == 0 {
				end = i
				boolend = true
			}
		}

		if boolend {
			str := ""
			for ii := start; ii <= end; ii++ {
				if len(lines[ii]) < 5 {
					continue
				}
				str = str + lines[ii] + "\n"
			}

			//过滤部分无效正文内容，以后扩展
			if strings.Contains(str, "Copyright") || strings.Contains(str, "版权所有") {
				continue
			}

			text = text + str
			boolstart = false
			boolend = false
		}

	}

	if start > end {
		size_1 := lenOfLines - 1
		str := ""
		for ii := start; ii <= size_1; ii++ {
			if len(lines[ii]) < 5 {
				continue
			}
			str = str + lines[ii] + "\n"
		}

		if !strings.Contains(str, "Copyright") && !strings.Contains(str, "版权所有") {
			text = text + str
		}
	}

	return text
}

func replaceBlank(src string) string {
	patterns := []string{
		" ", "",
		"\t", "",
	}
	replacer := strings.NewReplacer(patterns...)

	//src := "中(国)--中工(家伙)"
	strfmt := replacer.Replace(src)

	//	fmt.Println("\n  replacer.Replace old=", src)
	//	fmt.Println("  replacer.Replace new=", strfmt)

	return strfmt
}

package util

import (
	"github.com/huichen/sego"
	"strings"
)

type Segmenter struct {
	segmenter sego.Segmenter
}

func NewSegmenter() *Segmenter {
	client := new(Segmenter)
	// 载入词典
	dir := GetDir()
	client.segmenter.LoadDictionary(dir + "/../../huichen/sego/data/dictionary.txt")
	return client
}

// 分词
func (t *Segmenter) Segment(text string) []string {
	segments := t.segmenter.Segment([]byte(text))
	output := sego.SegmentsToSlice(segments, false)
	return filterWord(output)
}

// 过滤干扰词 空格，单字词
func filterWord(ws []string) []string {
	result := make([]string, 0)
	for _, w := range ws {
		if strings.Count(strings.TrimSpace(w), "") <= 2 {
			continue
		}
		result = append(result, w)
	}
	return result
}

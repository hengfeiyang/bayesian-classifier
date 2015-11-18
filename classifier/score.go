package classifier

import (
	"sort"
)

type Score struct {
	items []*ScoreItem
}

type ScoreItem struct {
	Category string  `json:"category"` // 分类名称
	Score    float64 `json:"score"`    // 概率值
}

func NewScores() *Score {
	t := new(Score)
	t.items = make([]*ScoreItem, 0)
	return t
}

func (t *Score) Append(category string, score float64) {
	t.items = append(t.items, &ScoreItem{category, score})
}

func (t *Score) GetSlice() []*ScoreItem {
	return t.items
}

func (t *Score) Top(n int) []*ScoreItem {
	t.Sort()
	if n == 0 || n > len(t.items) {
		n = len(t.items)
	}
	return t.items[0:n]
}

func (t *Score) Len() int {
	return len(t.items)
}

func (t *Score) Less(i, j int) bool {
	return t.items[i].Score > t.items[j].Score
}

func (t *Score) Swap(i, j int) {
	t.items[i], t.items[j] = t.items[j], t.items[i]
}

func (t *Score) Sort() {
	sort.Sort(t)
}

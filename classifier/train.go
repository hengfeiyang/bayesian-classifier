package classifier

import (
	"bayesian-classifier/util"
)

// 从指定的目录读取txt文件进行训练
func FileTrain(path string, classifier *Classifier) (int, error) {
	fs, err := util.ReadDir(path)
	if err != nil {
		return 0, err
	}
	i := 0
	for _, f := range fs {
		doc, err := util.ReadFile(path + "/" + f.Name())
		if err != nil {
			continue
		}
		category := getCategory(doc)
		if len(category) == 0 {
			continue
		}
		content := doc[len(category):]
		classifier.Training(string(content), string(category))
		i++
	}
	return i, nil
}

// 从文本文件中提取分类名称
// 第一行为分类名称
// 第二行为空行,分隔行
// 第三行向下为内容
func getCategory(doc []byte) []byte {
	category := make([]byte, 0)
	if len(doc) < 1 {
		return category
	}
	for i := 0; i < len(doc); i++ {
		if doc[i] != '\n' {
			category = append(category, doc[i])
		} else {
			break
		}
	}
	return category
}
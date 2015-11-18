// 贝叶斯分类器（Naive Bayesian classifier）支持中文文档解析训练和分类
package main

import (
	"fmt"
	"github.com/safeie/bayesian-classifier/classifier"
	"time"
)

// 如果开启HTTP，进程将不会立即结束，等待一个结束标志

const (
	DEFAULT_PROB      = 0.5            // 默认概率
	DEFUALT_WEIGHT    = 1.0            // 默认概率的权重，假定与一个单词相当
	DEBUG             = true           // 开启调试
	HTTP              = true           // 开启HTTP服务
	HTTP_PORT         = ":8812"        // HTTP端口
	STORAGE           = "file"         // 存储引擎，接受 file,redis，目前只支持file
	STORAGE_PATH      = "storage.data" // 文件存储引擎的存储路径
	STORAGE_FREQUENCY = "10"           // 自动存储的频率, 单位: 秒，0 表示不自动存储
)

func main() {
	// 分类器
	handler := classifier.NewClassifier(map[string]interface{}{
		"defaultProb":   DEFAULT_PROB,
		"defaultWeight": DEFUALT_WEIGHT,
		"debug":         DEBUG,
		"http":          HTTP,
		"httpPort":      HTTP_PORT,
		"storage": map[string]string{
			"adapter":   STORAGE,
			"path":      STORAGE_PATH,
			"frequency": STORAGE_FREQUENCY,
		},
	})

	// 训练
	handler.Training("这是一篇WEB开发的内容", "WEB")
	handler.Training("这是一篇Javascript的技巧", "WEB")
	handler.Training("这是一篇养生的内容", "WEB")
	handler.Training("这是一篇养生的内容2", "健康")
	handler.Training("这是一篇冬天养生食谱", "健康")
	handler.Training("坚持做运动就可以减肥", "测试")

	// 从txt文件进行训练
	classifier.FileTrain("./data", handler)

	// 获取训练数据
	testWord(handler, "养生", "WEB") // 测试已知分类
	testWord(handler, "养生", "XX")  // 测试未知分类
	testWord(handler, "养生", "")    // 查看所有分类
	testWord(handler, "不认识", "")   // 测试未知单词
	testWord(handler, "服务器", "")   // 测试未知单词

	// 分类测试
	testDoc(handler, "养生是什么分类")
	testDoc(handler, "API Go")
	testDoc(handler, "服务器")

	// 暂停
	time.Sleep(time.Second * 15)

	// 开启了HTTP服务，不能结束进程
	if HTTP {
		for {
			time.Sleep(time.Second)
		}
	}
}

// 辅助测试：测试单词的频率
func testWord(classifier *classifier.Classifier, word, category string) {
	score := classifier.Score(word, category)
	if category != "" {
		fmt.Printf("单词【%s】在分类【%s】中出现的概率为: \n", word, category)
	} else {
		fmt.Printf("单词【%s】在分类中出现的概率为: \n", word)
	}
	printScore(score)
}

// 辅助测试：测试文档的分类
func testDoc(classifier *classifier.Classifier, doc string) {
	score := classifier.Categorize(doc)
	fmt.Printf("测试文档归类于以下分类的概率为: \n")
	fmt.Println("--------------------------")
	fmt.Println(doc)
	fmt.Println("--------------------------")
	printScore(score)
}

// 辅助测试：输出
func printScore(scores []*classifier.ScoreItem) {
	if len(scores) == 0 {
		fmt.Println("未知单词 Orz！")
	}
	for k := range scores {
		fmt.Println(scores[k].Category, "\t", scores[k].Score)
	}
	fmt.Println(".")
}

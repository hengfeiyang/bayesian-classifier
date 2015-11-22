# 贝叶斯分类器（Naive Bayesian classifier）

通过一个spider抓取大量内容和分类交给分类器学习，使用朴素贝叶斯方法给未知内容分类。

提供一个web服务进行学习和测试。

将训练过的分类器保存起来可后续载入使用。

> 注意
> 
> 1. 对每篇学习过的内容进行HASH校验，避免大量重复内容影响分类器
> 2. 对于同一文档中重复出现的单词只计算一次

## 进度

1. 分类器构建完成
2. 基本接口可提供服务
3. 训练数据的保存和加载的存储器已完成
4. WEB Api已完成

## TODO

1. 抓个蜘蛛爬取大量数据构建测试
2. 性能优化，现在很占内存

## 问题

1. 发现现在的模型有问题，尝试分类文档时，由于每个单词的概率都是小于0的，所有单词的乘积是非常小的一个值，以至于被忽略成0了

## 使用

所有的例子都在 ``main.go``中演示了。请直接看源代码更实在。

依赖了一个分词器，要先安装一下

```
go get - u github.com/huichen/sego
```

使用就是 获取，导入，初始化，训练，测试

获取：

```
go get -u github.com/safeie/bayesian-classifier
```

测试：

```
package main

import (
	"github.com/safeie/bayesian-classifier/classifier"
	"fmt"
)

func main() {
	// 分类器
	handler := classifier.NewClassifier(map[string]interface{}{
		"defaultProb":   0.5,			// 默认概率
		"defaultWeight": 1.0,			// 默认概率的权重，假定与一个单词相当
		"debug":         false,			// 开启调试
		"http":          true,			// 开启HTTP服务
		"httpPort":      ":8812",		// HTTP服务端口
		"storage": map[string]string{
			"adapter":   "file",			// 存储引擎，接受 file,redis，目前只支持file
			"path":      "storage.data",	// 文件存储引擎的存储路径
			"frequency": "10",				// 自动存储的频率, 单位: 秒，0 表示不自动存储
		},
	})

	// 训练
	handler.Training("这是一篇WEB开发的内容", "web")
	handler.Training("这是一篇Javascript的技巧", "js")

	// 测试分类
	scores := handler.Categorize("你猜我说的这篇和开发有关的内容会是什么分类？")
	if len(scores) == 0 {
		fmt.Println("未知单词 Orz！")
	}
	for k := range scores {
		fmt.Println(scores[k].Category, "\t", scores[k].Score)
	}
}
```

## 接口

### 训练

给定文档和已知的分类进行训练

> 注意
>
> 1. 要对内容进行清洗，只要纯文本，去除HTML标识
> 2. 对内容进行分词
> 3. 去除干扰词（停止词），如：我们，他们，的，是


函数原型：

```
func Training(doc, category string) 
```

### 文本数据训练

方法：

```
classifier.FileTrain("./data", handler)
```

指定txt文件存储的路径，我放了几个在data目录中，传入分类器，会扫描目录下所有文件进行训练。

**格式要求：**

1. 文档为UTF-8编码
2. 文档第一行为 分类名称
3. 中间空一行分为分隔符
4. 从第三行开始识别为内容
5. 内容要求为纯文本格式，请自行过滤HTML标记


### 评分

获取给定单词在各分类中的概率，返回前10个按分值排序，如果指定分类参数，则只获取指定的分类。

函数原型：

```
func Score(word, category string) []*ScoreItem
```

### 分类

获取给定文档可能的分类，返回前10个按分值排序

函数原型：

```
func Categorize(doc string) []*ScoreItem
```

### 导出

将训练好的分类器数据到存储中，已备后续导入继续使用

函数原型：

```
func Export() error
```

### 导入

将存储的训练数据导入

函数原型：

```
func Import() error
```

##WEB API

接收GET/POST传参

训练：

```
http://localhost:8812/api/train?doc=文档&category=分类
```

获取单词统计：

```
http://localhost:8812/api/score?word=单词
```

分类：

```
http://localhost:8812/api/categorize?doc=文本内容
```

## 资料

### 朴素贝叶斯公式

```
P(A|B) = P(B|A) * P(A) / P(B)

P(category|document) = P(document|category) * P(category) / P(document)
P(category|document) = P(word1|category) * P(word2|category) ... * P(wordN|category)

```

1. P(category|document) 解释为：document在条件category下的概率，即文档属于该分类的概率
2. P(document|category) 解释为：category在条件document下的概率，即分类出现该文档的概率，又解释为，在该分类中的文档出现该文档所有单词的概率（分词后，该分类中文档出现单词N的概率，所有单词的概率的乘积）
3. P(category) 解释为：随机选择一篇文档属于该分类的概率，就是该分类的文档数除以文档总数
4. P(document) 解释为：``这篇文档出现的概率？``
5. P(word|category) 解释为：该分类的文档中出现这个单词的概率

    至于P(document)，我们也可以计算它，但这将会是一项不必要的工作。请记住，我们不会将这一计算当做真实的概率值。相反，我们会分别计算每个分类的概率，然后对所有的计算结果进行比较。由于不论是哪个分类，P(Document)的值都是一样的，其对结果所产生的影响也是完全一样的，因为我们完全可以忽略这一项。


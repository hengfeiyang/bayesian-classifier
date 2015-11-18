package classifier

import (
	"fmt"
	"log"
	"net/http"
)

type Http struct {
	port       string
	classifier *Classifier
}

func NewHttp(port string, classifier *Classifier) *Http {
	t := new(Http)
	t.port = port
	t.classifier = classifier
	return t
}

func (t *Http) Start() {
	http.HandleFunc("/", t.handleIndex)
	http.HandleFunc("/api/train", t.handleTrain)
	http.HandleFunc("/api/score", t.handleScore)
	http.HandleFunc("/api/categorize", t.handleGategorize)
	log.Fatal(http.ListenAndServe(t.port, nil))
}

func (t *Http) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is index")
}

// 训练接口
func (t *Http) handleTrain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", t.classifier.Data.Categorys)
}

// 获取单词分类接口
func (t *Http) handleScore(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is score")
}

// 分类接口
func (t *Http) handleGategorize(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is categorize")
}

package classifier

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/safeie/bayesian-classifier/util"
)

// HTTP 提供了HTTP接口的结构体
type HTTP struct {
	port       string
	tplDir     string
	assetsDir  string
	classifier *Classifier
}

type result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewHTTP create a HTTP
func NewHTTP(port string, classifier *Classifier) *HTTP {
	t := new(HTTP)
	t.port = port
	rootDir := util.GetDir()
	t.tplDir = rootDir + "/html"
	t.assetsDir = rootDir + "/assets"
	t.classifier = classifier
	return t
}

// Start the http service
func (t *HTTP) Start() {
	http.HandleFunc("/", t.handleIndex)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(t.assetsDir))))
	http.HandleFunc("/api/train", t.handleTrain)
	http.HandleFunc("/api/score", t.handleScore)
	http.HandleFunc("/api/categorize", t.handleGategorize)
	log.Fatal(http.ListenAndServe(t.port, nil))
}

func (t *HTTP) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(t.tplDir + "/index.html")
	if err != nil {
		fmt.Fprintln(w, output(1, err.Error(), nil))
		return
	}
	tpl.Execute(w, nil)
}

// 训练接口
func (t *HTTP) handleTrain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	doc := r.FormValue("doc")
	category := r.FormValue("category")
	t.classifier.Training(doc, category)
	fmt.Fprintf(w, "%s", output(0, "", nil))
}

// 获取单词分类接口
func (t *HTTP) handleScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	word := r.FormValue("word")
	category := r.FormValue("category")
	scores := t.classifier.Score(word, category)
	fmt.Fprintf(w, "%s", output(0, "", scores))
}

// 分类接口
func (t *HTTP) handleGategorize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	doc := r.FormValue("doc")
	scores := t.classifier.Categorize(doc)
	fmt.Fprintf(w, "%s", output(0, "", scores))
}

// 输出返回结果JSON
func output(code int, message string, data interface{}) []byte {
	res := new(result)
	res.Code = code
	res.Message = message
	res.Data = data
	jsonStr, err := json.Marshal(res)
	if err != nil {
		jsonStr = []byte("{\"code\":1, \"message\":\"" + err.Error() + "\"}")
	}
	return jsonStr
}

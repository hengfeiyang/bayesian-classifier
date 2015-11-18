package classifier

import (
	"encoding/json"
	"fmt"
	//"github.com/safeie/bayesian-classifier/util"
	"log"
	"net/http"
)

type Http struct {
	port       string
	classifier *Classifier
}

type result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHttp(port string, classifier *Classifier) *Http {
	t := new(Http)
	t.port = port
	t.classifier = classifier
	return t
}

func (t *Http) Start() {
	//dir := util.GetDir()
	http.HandleFunc("/", t.handleIndex)
	//http.Handle("/assets", http.FileServer(dir+"/assets"))
	http.HandleFunc("/api/train", t.handleTrain)
	http.HandleFunc("/api/score", t.handleScore)
	http.HandleFunc("/api/categorize", t.handleGategorize)
	log.Fatal(http.ListenAndServe(t.port, nil))
}

func (t *Http) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "this is index")
}

// 训练接口
func (t *Http) handleTrain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	doc := r.FormValue("doc")
	category := r.FormValue("category")
	t.classifier.Training(doc, category)
	fmt.Fprintf(w, "%s", output(0, "", nil))
}

// 获取单词分类接口
func (t *Http) handleScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	word := r.FormValue("word")
	category := r.FormValue("category")
	scores := t.classifier.Score(word, category)
	fmt.Fprintf(w, "%s", output(0, "", scores))
}

// 分类接口
func (t *Http) handleGategorize(w http.ResponseWriter, r *http.Request) {
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

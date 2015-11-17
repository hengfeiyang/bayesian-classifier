package storage

import (
	"bayesian-classifier/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

type FileStorage struct {
	path string
}

func NewFileStorage(path string) (*FileStorage, error) {
	t := new(FileStorage)
	// 修正路径
	if path[0] != '/' {
		rootDir, _ := util.GetDir()
		path = rootDir + "/" + path
	}
	t.path = path
	log.Println("数据存储文件：", t.path)

	// 判断存储路径是否存在
	ok, err := util.IsExist(filepath.Base(path))
	if ok == false || err != nil {
		return nil, err
	}
	// 判断路径是否可写
	ok, err = util.IsWritable(path)
	if ok == false || err != nil {
		return nil, err
	}

	return t, nil
}

// 使用JSON格式存储
func (t *FileStorage) Save(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(t.path, jsonData, 0666)
}

// 使用JSON格式加载
func (t *FileStorage) Load(target interface{}) error {
	ok, err := util.IsExist(t.path)
	if err != nil {
		return err
	}
	// 不存在存储的数据, 跳过加载
	if ok == false {
		return nil
	}
	// 加载JSON数据
	stData, err := util.ReadFile(t.path)
	if err != nil {
		return err
	}
	return json.Unmarshal(stData, target)
}

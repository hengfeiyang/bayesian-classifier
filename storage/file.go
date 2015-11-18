package storage

import (
	"bayesian-classifier/util"
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type FileStorage struct {
	path string // 存储路径
	sign string // 数据签名，用来判断是否是否较加载时修改过
}

func NewFileStorage(path string) (*FileStorage, error) {
	t := new(FileStorage)
	// 修正路径
	if path[0] != '/' {
		rootDir := util.GetDir()
		path = rootDir + "/" + path
	}
	t.path = path

	// 判断存储路径是否存在
	ok := util.IsExist(filepath.Dir(path))
	if ok == false {
		return nil, errors.New("存储路径不存在: " + filepath.Dir(path))
	}
	// 判断路径是否可写
	ok = util.IsWritable(filepath.Dir(path))
	if ok == false {
		return nil, errors.New("存储路径不可写: " + filepath.Dir(path))
	}

	return t, nil
}

// 使用JSON格式存储
func (t *FileStorage) Save(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	sign := util.MD5(string(jsonData))
	if t.sign == "" {
		t.sign = sign
	} else {
		// 签名相同，数据没变化
		if t.sign == sign {
			return nil
		} else {
			t.sign = sign
		}
	}
	return ioutil.WriteFile(t.path, jsonData, 0666)
}

// 使用JSON格式加载
func (t *FileStorage) Load(target interface{}) error {
	// 不存在存储的数据, 跳过加载
	ok := util.IsExist(t.path)
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

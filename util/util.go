package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// 执行系统命令并返回结果
func Command(pro string, argv []string, baseDir string) ([]byte, error) {
	cmd := exec.Command(pro, argv...)
	// 设置命令运行时目录
	if baseDir != "" {
		cmd.Dir = baseDir
	}
	res, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 获取程序运行的目录
func GetDir() string {
	path, err := filepath.Abs(os.Args[0])
	if err != nil {
		return ""
	}
	return filepath.Dir(path)
}

// 判断一个文件或目录是否存在
// 存在时返回 true, 不存在返回 false
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	// Check if error is "no such file or directory"
	if _, ok := err.(*os.PathError); ok {
		return false
	}
	return false
}

// 判断一个文件或目录是否有写入权限
// 可写时返回 true, 不可写返回 false
func IsWritable(path string) bool {
	err := syscall.Access(path, syscall.O_RDWR)
	if err == nil {
		return true
	}
	// Check if error is "no such file or directory"
	if _, ok := err.(*os.PathError); ok {
		return false
	}
	return false
}

// 读取一个文件夹返回文件列表
func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 获取一个文件的文件后缀名
func GetExt(filename string) string {
	info := strings.Split(filename, ".")
	if len(info) < 2 {
		return ""
	}
	return info[len(info)-1]
}

// 复制文件，仅文件，不支持目录
func CopyFile(s, d string) error {
	// 坑爹啊，要先删除是不是link
	linfo, err := os.Readlink(s)
	if err == nil || len(linfo) > 0 {
		// 这货是link，创建link吧
		return os.Symlink(linfo, d)
	}
	// 不是link，创建文件
	sf, err := os.Open(s)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(d)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	return err
}

// 读取文件返回内容
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

// 计算一个字符串的MD5
func MD5(str string) string {
	hexStr := md5.Sum([]byte(str))
	return hex.EncodeToString(hexStr[:])
}

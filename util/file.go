package util

import (
	"bufio"
	"go/ast"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func LoadAllGoFile(path string) ([]string, error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0)
	err = getAllGoFile(path, dir, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getAllGoFile(parent string, dir []fs.FileInfo, res *[]string) error {
	for _, file := range dir {
		absolutePath := path.Join(parent, file.Name())
		if file.IsDir() {
			readDir, err := ioutil.ReadDir(absolutePath)
			if err != nil {
				return err
			}
			err = getAllGoFile(absolutePath, readDir, res)
			if err != nil {
				return err
			}
		} else {
			if strings.HasSuffix(file.Name(), ".go") {
				*res = append(*res, absolutePath)
			}
		}
	}
	return nil
}

func ConvertConfigMap(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = ConvertConfigMap(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = ConvertConfigMap(v)
		}
	}
	return i
}

func GetFilePackage(file *ast.File) string {
	return file.Name.Name
}
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func CreateFile(path string, code []byte) error {
	length := len(path)
	pos := length
	for pos = length - 1; pos >= 0; pos-- {
		if path[pos] == '/' {
			break
		}
	}
	if path == "" {
		return nil
	}
	prefix := path[:pos]
	if !isExist(prefix) {
		err := os.MkdirAll(prefix, os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	writer := bufio.NewWriter(file)
	if err != nil {
		return err
	}
	_, err = writer.Write(code)
	writer.Flush()
	return err
}

func Clean(gc []string) error {
	for _, path := range gc {
		os.Remove(path)
	}
	return nil
}

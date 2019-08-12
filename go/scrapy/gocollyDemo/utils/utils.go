/**
 * @Author: FB
 * @Description: 
 * @File:  utils.go
 * @Version: 1.0.0
 * @Date: 2019/7/3 14:11
 */

package CollyUtils

import (
	"fmt"
	"github.com/mahonia"
	"os"
	"bufio"
	"net/http"
	"io/ioutil"
	"io"
	"bytes"
	"strings"
	"errors"
	)

func ShowVals(vals ...interface{}) {
	for _, v := range vals {
		fmt.Println(v)
	}
}

func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)

	return result

}

func AppendToFile(file, str string) (err error) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	num, err := writer.WriteString(str)
	fmt.Println("num=", num)
	err = writer.Flush()

	return

}

func DownloadImage(url string, save_path_dir string) (err error) {

	name := GetImageNameFromUrl(url)
	fmt.Println(name)

	filename := save_path_dir + name

	//判定目录是否存在，不存在，则创建
	//dir, _ := filepath.Split(path)
	DirIsExist(save_path_dir)

	//判定是否已经存在
	err = IsFileExisted(filename)

	// 说明文件已经存在，不需要下载
	if err == nil {
		fmt.Println("file has existed,file=", filename)
		return
	}

	out, err := os.Create(filename)
	//log.Println("DownloadImage getImg001", err)
	defer out.Close()
	if err != nil {
		return
	}

	resp, err := http.Get(url)
	//log.Println("DownloadImage getImg002", err)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	pix, err := ioutil.ReadAll(resp.Body)
	//log.Println("DownloadImage getImg003", err)
	if err != nil {
		return
	}
	_, err = io.Copy(out, bytes.NewReader(pix))
	//log.Println("DownloadImage getImg004", err)
	return

}

func GetImageNameFromUrl(url string) (name string) {
	path := strings.Split(url, "/")
	if len(path) >= 1 {
		name = path[len(path)-1]
	}
	return
}

// 只判断文件是否存在
func IsFileExisted(path string) (err error) {
	_, err = os.Stat(path)
	if err != nil {
		return err
	}
	return err
}

// 判断目录是否存在,不存在则新建一个
func DirIsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		ok, err := CreateDir(path)
		if !ok {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, err
}

// 新建目录
func CreateDir(path string) (bool, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		errInfo := errors.New(fmt.Sprintf("Mkdir %s failed! error= %v\n", path, err))
		return false, errInfo
	} else {
		return true, nil
	}
}

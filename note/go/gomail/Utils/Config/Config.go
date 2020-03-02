/**
 * @Author: FB
 * @Description: 
 * @File:  Config.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 15:18
 */

package Config

import (
	"strings"
	"os"
	"bufio"
	"io"
	"gomail/Utils/Common"
)

const (
	// Windows
	WINDOWS_CONFIG_PATH = `D:\Go-Work\src\gomail\config.txt`
	WINDOWS_LOCAL_CONFIG_PATH   = `D:\Go-Work\src\gomail\local_config.txt`
	// Linux
	LINUX_CONFIG_PATH = `/usr/local/Go-Work/src/gomail/config.txt`
	LINUX_LOCAL_CONFIG_PATH   = `/usr/local/Go-Work/src/gomail/local_config.txt`
)

func GetValue(key string) (v string) {
	config := getConfig()
	v = config[key]
	return
}

func getConfig() (config map[string]string) {
	config = make(map[string]string)

	array := getConfigFile(1)
	for i := 0; i < len(array); i++ {
		if strings.EqualFold(array[i], "") {
			break
		}
		kv := strings.Split(array[i], "=")
		config[kv[0]] = kv[1]
	}

	v := config["ip"]

	switch v {
	case "168":
		{
			array = getConfigFile(2)
		}
	default:
		return
	}

	for i := 0; i < len(array); i++ {
		if strings.EqualFold(array[i], "") {
			return
		}
		kv := strings.Split(array[i], "=")
		config[kv[0]] = kv[1]
	}

	return
}

func getConfigFile(flag int) (config []string) {
	var f *os.File
	var err error

	if CommonUtils.IsWindowsPlatform() {
		switch flag {
		case 1:
			f, err = os.Open(WINDOWS_CONFIG_PATH)
		case 2:
			f, err = os.Open(WINDOWS_LOCAL_CONFIG_PATH)
		}
	} else {
		switch flag {
		case 1:
			f, err = os.Open(LINUX_CONFIG_PATH)
		case 2:
			f, err = os.Open(LINUX_LOCAL_CONFIG_PATH)
		}
	}
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil && err != io.EOF {
			break
		}

		if err == io.EOF && len(line) == 0 {
			return
		}

		// 过滤注释和换行
		if strings.EqualFold(line[:2], "//") || len(line) == 2 {
			continue
		}

		if err == io.EOF && len(line) > 0 {
			config = append(config, line)
		} else {
			config = append(config, line[:len(line)-2])
		}
	}

	return
}

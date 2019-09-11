/**
 * @Author: FB
 * @Description: 
 * @File:  CommonUtils.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 14:30
 */

package CommonUtils

import (
	"encoding/json"
	"strings"
	"runtime"
	"fmt"
	"gomail/Utils/CommonVar"
)

func GetNormalRespJson(status int, errcode interface{}, data map[string]interface{}) string {
	resp_map := make(map[string]interface{})
	resp_map["status"] = status
	resp_map["errcode"] = errcode
	resp_map["data"] = data

	b, _ := json.Marshal(resp_map)
	fmt.Println("GetNormalRespJson=", string(b), "errmsg=", CommonVar.CodeStatusMap[errcode.(int)])
	return string(b)
}

func GetListRespJson(status int, errcode interface{}, data []map[string]interface{}, total int) string {
	resp_map := make(map[string]interface{})
	resp_map["status"] = status
	resp_map["errcode"] = errcode
	resp_map["total"] = total
	resp_map["data"] = data

	b, _ := json.Marshal(resp_map)
	fmt.Println("GetListRespJson=", string(b), "errmsg=", CommonVar.CodeStatusMap[errcode.(int)])
	return string(b)
}

func IsLinuxPlatform() bool {
	return strings.EqualFold(runtime.GOOS, "linux")
}

func IsWindowsPlatform() bool {
	return strings.EqualFold(runtime.GOOS, "windows")
}

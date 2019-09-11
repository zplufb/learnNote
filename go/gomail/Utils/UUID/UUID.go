/**
 * @Author: FB
 * @Description: 
 * @File:  UUID.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 16:33
 */

package UUID

import (
	"github.com/xid"
	"math/rand"
	"time"
	"fmt"
	"strconv"
	"math"
	"encoding/base64"
)

// 获取一个XID
func GetXID() (id string) {
	guid := xid.New()
	id = guid.String()

	return
}


func GetRandCode(number int) string {

	if number > 8 || number < 1{
		return "-1"
	}
	//number=4,6,8
	num_str := strconv.Itoa(number)
	maxNum := int32(math.Pow10(number))

	randNum := fmt.Sprintf("%0"+num_str+"v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(maxNum))

	return randNum
}

// base64解码
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// base64编码
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

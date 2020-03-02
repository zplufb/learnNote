/**
 * @Author: FB
 * @Description: 
 * @File:  GlobalVar.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 17:51
 */

package CommonVar

import "gomail/Model"

var BasicSettingMap map[int]Model.BasicSetting
var MailSettingMap map[int]Model.MailSetting

const (
	TYPE_BS_TIME    = 0
	TYPE_BS_SUBJECT = 1
	TYPE_BS_BODY    = 2

	TYPE_MS_ADMIN   = 0
	TYPE_BS_SUPPORT = 1
)

/**
 * @Author: FB
 * @Description: 
 * @File:  InitServer.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 17:39
 */

package InitServer

import (
	"gomail/Server/BasicSettingServer"
	"gomail/Server/MailSettingServer"
	"log"
	"gomail/Utils/CommonVar"
	"gomail/Model"
)

func InitVar() {

	var BasicSettingMap = make(map[int]Model.BasicSetting)
	var MailSettingMap = make(map[int]Model.MailSetting)

	ok, errcode, listBS, total := BasicSettingServer.GetBasicSettingListInit(10, 0)
	if ok && total > 0 {
		for i, v := range listBS {
			BasicSettingMap[i] = v
		}
	} else {
		log.Println("errcode=", errcode)
	}

	CommonVar.BasicSettingMap = BasicSettingMap

	ok, errcode, listMS, total := MailSettingServer.GetMailSettingListInit(10, 0)
	if ok && total > 0 {
		for i, v := range listMS {
			MailSettingMap[i] = v
		}
	} else {
		log.Println("errcode=", errcode)
	}

	CommonVar.MailSettingMap = MailSettingMap

}

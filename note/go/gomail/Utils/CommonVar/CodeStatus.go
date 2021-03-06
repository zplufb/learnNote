/**
 * @Author: FB
 * @Description: 
 * @File:  CodeStatusUtils.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 14:43
 */

package CommonVar

//1000是成功相关（模块（1001,1100,1200,1300分别是验证码，发送邮件，邮箱账号，基本配置），2000是数据库等相关，，4000是失败
const (
	SUCCEE                      = 1000
	FAIL                        = 4000
	REQUEST_PARAMETER_INCORRECT = 4001
	INVALID_MAIL_FORMAT         = 4003

	VERIFY_CODE_CREATE_SUCCEE     = 1001
	VERIFY_CODE_CREATE_FAIL       = 1011
	VERIFY_CODE_GET_RECORD_SUCCEE = 1002
	VERIFY_CODE_GET_RECORD_FAIL   = 1012
	VERIFY_CODE_CHECK_SUCCEE      = 1003
	VERIFY_CODE_CHECK_FAIL        = 1013
	VERIFY_CODE_SEND_SUCCEE       = 1004
	VERIFY_CODE_SEND_FAIL         = 1014
	VERIFY_CODE_TIMEOUT           = 1015

	SEND_MAIL_CREATE_SUCCEE = 1100
	SEND_MAIL_CREATE_FAIL   = 1110
	SEND_MAIL_LIST_SUCCEE   = 1101
	SEND_MAIL_LIST_FAIL     = 1111
	SEND_MAIL_SEND_SUCCEE   = 1102
	SEND_MAIL_SEND_FAIL     = 1112

	MAIL_SETTING_CREATE_SUCCEE = 1200
	MAIL_SETTING_CREATE_FAIL   = 1210
	MAIL_SETTING_LIST_SUCCEE   = 1201
	MAIL_SETTING_LIST_FAIL     = 1211

	BASIC_SETTING_CREATE_SUCCEE     = 1300
	BASIC_SETTING_CREATE_FAIL       = 1310
	BASIC_SETTING_LIST_SUCCEE       = 1301
	BASIC_SETTING_LIST_FAIL         = 1311
	BASIC_SETTING_GET_RECORD_SUCCEE = 1302
	BASIC_SETTING_GET_RECORD_FAIL   = 1312

	DB_CONNECT_SUCCEE = 2000
	DB_CONNECT_FAIL   = 2010
)

var CodeStatusMap = map[int]string{

	SUCCEE:                      "成功",
	FAIL:                        "失败",
	REQUEST_PARAMETER_INCORRECT: "请求参数错误！",
	INVALID_MAIL_FORMAT:         "无效邮箱格式",

	//1001
	VERIFY_CODE_CREATE_SUCCEE:     "验证码创建成功",
	VERIFY_CODE_CREATE_FAIL:       "验证码创建失败",
	VERIFY_CODE_GET_RECORD_SUCCEE: "验证码获取成功",
	VERIFY_CODE_GET_RECORD_FAIL:   "验证码获取失败",
	VERIFY_CODE_CHECK_SUCCEE:      "验证码验证成功",
	VERIFY_CODE_CHECK_FAIL:        "验证码验证失败",
	VERIFY_CODE_SEND_SUCCEE:       "验证码发送成功",
	VERIFY_CODE_SEND_FAIL:         "验证码发送失败",
	VERIFY_CODE_TIMEOUT:           "验证码已过期",
	//1100
	SEND_MAIL_CREATE_SUCCEE: "邮件发送记录创建成功",
	SEND_MAIL_CREATE_FAIL:   "邮件发送记录创建失败",
	SEND_MAIL_LIST_SUCCEE:   "邮件发送记录列表获取成功",
	SEND_MAIL_LIST_FAIL:     "邮件发送记录列表获取失败",
	SEND_MAIL_SEND_SUCCEE:   "邮件发送成功",
	SEND_MAIL_SEND_FAIL:     "邮件发送失败",
	//1200
	MAIL_SETTING_CREATE_SUCCEE: "邮箱账号创建成功",
	MAIL_SETTING_CREATE_FAIL:   "邮箱账号创建失败",
	MAIL_SETTING_LIST_SUCCEE:   "邮箱账号列表获取成功",
	MAIL_SETTING_LIST_FAIL:     "邮箱账号列表获取失败",
	//1300
	BASIC_SETTING_CREATE_SUCCEE:     "基本配置记录创建成功",
	BASIC_SETTING_CREATE_FAIL:       "基本配置记录创建失败",
	BASIC_SETTING_LIST_SUCCEE:       "基本配置记录列表获取成功",
	BASIC_SETTING_LIST_FAIL:         "基本配置记录列表获取失败",
	BASIC_SETTING_GET_RECORD_SUCCEE: "基本配置记录获取成功",
	BASIC_SETTING_GET_RECORD_FAIL:   "基本配置记录获取失败",
	//2000
	DB_CONNECT_SUCCEE: "数据库连接成功",
	DB_CONNECT_FAIL:   "数据库连接失败",
}

package test

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"testing"
	"time"

	"github.com/VINDA-98/Fasthop/global"
)

var (
	SmtpMailUser     = global.App.Config.Smtp.User
	SmtpMailPwd      = global.App.Config.Smtp.Password
	SmtpMailHost     = global.App.Config.Smtp.Host
	SmtpMailNickname = global.App.Config.Smtp.Nickname
	SmtpMailPort     = global.App.Config.Smtp.Port
	SmtpMailTarget   = global.App.Config.Smtp.Target
)

type Req struct {
	Stories []Story `json:"stroies"`
}
type Story struct {
	Uuid   string `json:"uuid"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Number string `json:"number"`
}

func TestSendMail(t *testing.T) {
	err := SendMail()
	if err != nil {
		return
	}
}

func SendMail() (err error) {

	addressArray := strings.Split(SmtpMailTarget, ",") //目标提醒用户邮件地址

	//两天后的日期
	endStr := time.Now().AddDate(0, 0, 2)
	week := ""
	switch endStr.Weekday() {
	case 0:
		week = "星期天"
		break
	case 1:
		week = "星期一"
		break
	case 2:
		week = "星期二"
		break
	case 3:
		week = "星期三"
		break
	case 4:
		week = "星期四"
		break
	case 5:
		week = "星期五"
		break
	case 6:
		week = "星期六"
		break
	default:
		week = "星期五"
		break
	}

	Req := &Req{}
	story1 := Story{
		Uuid:   "5zmUJGm8NKhe68qY",
		Title:  "作为项目经理，我希望部分工作项属性可以直接引用项目属性的值，无需每个工作项手工维护数据，提高工作效率",
		Desc:   "1、现状：部分工作项属性，需引用项目属性的值，如“客户”、“合同”，统计分析时，需根据这些属性统计工作项数量或其他数据等，目前是工作项属性中建立自定义属性，在创建工作项时手工再维护一次项目属性中已维护的值，工作量较大\n2、需求：2.1 工作项属性的值，可以直接引用项目属性的值，无需再手工录入\n2.2 希望这一类属性可以直接显示在详情页的【基础信息】模块\n2.3 此类配置可通过项目模板复制，无需每个项目配置一次\n2.4 切换时，历史工作项也希望能引用对应的值，不只是新建工作项可引用",
		Number: "57",
	}
	story2 := Story{
		Uuid:   "5zmUJGm8NKhe68qG",
		Title:  "作为项目经理，我希望部分工作项属性可以直接引用项目属性的值，无需每个工作项手工维护数据，提高工作效率aaaaaaa",
		Desc:   "1、现状：部分工作项属性，需引用项目属性的值，如“客户”、“合同”，统计分析时，需根据这些属性统计工作项数量或其他数据等，目前是工作项属性中建立自定义属性，在创建工作项时手工再维护一次项目属性中已维护的值，工作量较大\n2、需求：2.1 工作项属性的值，可以直接引用项目属性的值，无需再手工录入\n2.2 希望这一类属性可以直接显示在详情页的【基础信息】模块\n2.3 此类配置可通过项目模板复制，无需每个项目配置一次\n2.4 切换时，历史工作项也希望能引用对应的值，不只是新建工作项可引用",
		Number: "58",
	}

	Req.Stories = append(Req.Stories, story1)
	Req.Stories = append(Req.Stories, story2)
	var numbers []string

	internalBody := ""
	for _, story := range Req.Stories {
		numbers = append(numbers, story.Number) //增加编号
		internalBody += fmt.Sprintf(`content %s %s %s`, story.Number, story.Title, story.Uuid) + "\r\n"
	}

	log.Println("body:", internalBody)

	subject := fmt.Sprintf("%s号需求报价 %s", numbers, week)

	// 通常身份应该是空字符串，填充用户名.
	auth := smtp.PlainAuth("", SmtpMailUser, SmtpMailPwd, SmtpMailHost)
	contentType := "Content-Type: text/html; charset=UTF-8"
	for _, v := range addressArray {
		s := fmt.Sprintf(
			"To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n",
			v,
			SmtpMailNickname,
			SmtpMailUser,
			subject,
			contentType,
		)
		msg := []byte(s)
		addr := fmt.Sprintf("%s:%s", SmtpMailHost, SmtpMailPort)
		err = smtp.SendMail(addr, auth, SmtpMailUser, []string{v}, msg)
		if err != nil {
			return err
		}
	}
	return
}

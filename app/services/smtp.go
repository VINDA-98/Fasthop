package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/VINDA-98/Fasthop/utils"

	"github.com/VINDA-98/Fasthop/app/models"
	"github.com/VINDA-98/Fasthop/global"
)

// @Title  services
// @Description  MyGO
// @Author  WeiDa  2023/5/6 15:30
// @Update  WeiDa  2023/5/6 15:30

type smtpService struct {
}

var SmtpService = new(smtpService)

func (s *smtpService) SendMail(stories []models.Story) (err error) {
	if len(stories) == 0 {
		return
	}
	if global.App.Config.Smtp.IsEmpty() {
		return
	}

	body := `

		tips:邮件推荐使用网易灵犀办公打开
	`
	subject := fmt.Sprintf("这是主题")

	// 发送邮件
	addressArray := strings.Split(global.App.Config.Smtp.Target, ",") //目标提醒用户邮件地址
	log.Println("正在发送邮件到:", addressArray)
	client := utils.NewEmailClient(
		global.App.Config.Smtp.User,
		global.App.Config.Smtp.Password,
		global.App.Config.Smtp.Nickname,
		global.App.Config.Smtp.Host,
		global.App.Config.Smtp.Port,
		true,
	)
	if err := client.SendEmail(addressArray, subject, body); err != nil {
		log.Println("发送邮件失败:", err)
	}
	return
}

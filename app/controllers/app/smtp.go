package app

// @Title  app
// @Description  MyGO
// @Author  WeiDa  2023/5/6 16:17
// @Update  WeiDa  2023/5/6 16:17

import (
	"fmt"
	"github.com/VINDA-98/Fasthop/global"
	"github.com/VINDA-98/Fasthop/utils"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func SendMailTest(c *gin.Context) {

	body := `

		tips:邮件推荐使用网易灵犀办公打开
	`

	subject := fmt.Sprintf("这是测试的主题")

	addressArray := strings.Split("test@qq.com", ",") //目标提醒用户邮件地址

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

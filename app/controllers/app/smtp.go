package app

// @Title  app
// @Description  MyGO
// @Author  WeiDa  2023/5/6 16:17
// @Update  WeiDa  2023/5/6 16:17

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/VINDA-98/Fasthop/global"
	"github.com/VINDA-98/Fasthop/utils"

	"github.com/VINDA-98/Fasthop/app/services"

	"github.com/VINDA-98/Fasthop/app/models"

	"github.com/gin-gonic/gin"
)

func SendMailTest(c *gin.Context) {
	startStr := time.Now().Format("2006年01月02日")
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

	var stories []models.Story

	story1 := models.Story{
		UUID:   "5zmUJGm8NKhe68qY",
		Title:  "作为项目经理，我希望部分工作项属性可以直接引用项目属性的值，无需每个工作项手工维护数据，提高工作效率",
		Desc:   "1、现状：部分工作项属性，需引用项目属性的值，如“客户”、“合同”，统计分析时，需根据这些属性统计工作项数量或其他数据等，目前是工作项属性中建立自定义属性，在创建工作项时手工再维护一次项目属性中已维护的值，工作量较大\n2、需求：2.1 工作项属性的值，可以直接引用项目属性的值，无需再手工录入\n2.2 希望这一类属性可以直接显示在详情页的【基础信息】模块\n2.3 此类配置可通过项目模板复制，无需每个项目配置一次\n2.4 切换时，历史工作项也希望能引用对应的值，不只是新建工作项可引用",
		Number: "88",
	}
	story2 := models.Story{
		UUID:   "5zmUJGm8NKhe68qG",
		Title:  "作为项目经理，我希望部分工作项属性可以直接引用项目属性的值，无需每个工作项手工维护数据，提高工作效率aaaaaaa",
		Desc:   "1、现状：部分工作项属性，需引用项目属性的值，如“客户”、“合同”，统计分析时，需根据这些属性统计工作项数量或其他数据等，目前是工作项属性中建立自定义属性，在创建工作项时手工再维护一次项目属性中已维护的值，工作量较大\n2、需求：2.1 工作项属性的值，可以直接引用项目属性的值，无需再手工录入\n2.2 希望这一类属性可以直接显示在详情页的【基础信息】模块\n2.3 此类配置可通过项目模板复制，无需每个项目配置一次\n2.4 切换时，历史工作项也希望能引用对应的值，不只是新建工作项可引用",
		Number: "89",
	}

	stories = append(stories, story1)
	stories = append(stories, story2)
	var numbers []string

	body := ""
	for _, story := range stories {
		numbers = append(numbers, story.Number) //增加编号
		Group1WikiUUID := services.SmtpService.GetWikiUUID(fmt.Sprintf("#%s ONES交付一组-报价", story.Number))
		Group1WikiLink := fmt.Sprintf("%s/wiki/#/team/RDjYMhKq/space/%s/page/%s",
			global.App.Config.ONES.Base,
			global.App.Config.ONES.SpaceUUID,
			Group1WikiUUID,
		)
		Group2WikiUUID := services.SmtpService.GetWikiUUID(fmt.Sprintf("#%s ONES交付三组-报价", story.Number))
		Group2WikiLink := fmt.Sprintf("%s/wiki/#/team/RDjYMhKq/space/%s/page/%s",
			global.App.Config.ONES.Base,
			global.App.Config.ONES.SpaceUUID,
			Group2WikiUUID,
		)

		body += fmt.Sprintf(`
			%s号需求
			%s。
			%s/project/#/team/RDjYMhKq/task/%s
			
			ONES交付一组链接：
			%s

			ONES交付三组链接：
			%s
			-------------------------------------------------------
		`,
			story.Number,
			story.Title,
			global.App.Config.ONES.Base,
			story.UUID,
			Group1WikiLink,
			Group2WikiLink,
		) + "\n"
	}
	body += fmt.Sprintf(`以上需求请在两天内（截止时间：【%s（%s）】）完成，谢谢！


			`, endStr.Format("2006年01月02日"), week)

	body += fmt.Sprintf(`
			光恒  \n
			您好，【%s】新增%d个需求%s号需求，状态已更新为下单中：  \n
			其中%s号需求请在【%s（%s）】完成报价。 \n
			需求列表：  \n
			https://partner.ones.ai/project/#/team/GzPY6Hs8/project/UJpAKvZyVTcsXCle/component/A7s6KQFY/view/2DoUuE6J  \n

			揽月  \n
			您好，【%s】新增%d个需求%s号需求，状态已更新为下单中：   \n
			其中%s号需求请在【%s（%s）】完成报价。  \n
			需求列表：  \n
			https://partner.ones.ai/project/#/team/P4fGixEe/project/UJpAKvZytcIhwrZ7/component/TRinZhfP/view/MeodL2Ys  \n

			锦高  \n
			您好，【%s】新增%d个需求%s号需求，状态已更新为下单中：   \n
			其中%s号需求请在【%s（%s）】完成报价。 \n 
			需求列表： \n 
			https://partner.ones.ai/project/#/team/V1FJr5kd/project/UJpAKvZyOo3ZXLng/component/JxMPKM9a/view/Wj7DcqfY  \n
			
	`,
		startStr, len(numbers), numbers, numbers, endStr.Format("2006年01月02日"), week,
		startStr, len(numbers), numbers, numbers, endStr.Format("2006年01月02日"), week,
		startStr, len(numbers), numbers, numbers, endStr.Format("2006年01月02日"), week,
	)
	body += `

		tips:邮件推荐使用网易灵犀办公打开
	`

	subject := fmt.Sprintf("这是测试的 => %s号需求报价", numbers)

	// 发送邮件 @todo 测试只发开发者
	addressArray := strings.Split("weida@ones.cn", ",") //目标提醒用户邮件地址

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

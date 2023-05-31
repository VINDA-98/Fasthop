package services

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"

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

	var numbers []string
	start := time.Now()
	end := time.Now().AddDate(0, 0, 2)
	startDate := start.Format("2006年01月02日")
	//结束日期
	week := ""
	switch start.Weekday() {
	case 0:
		week = "下周二"
		break
	case 1:
		week = "本周三"
		break
	case 2:
		week = "本周四"
		break
	case 3:
		week = "本周五"
		break
	case 4:
		week = "下周一" //星期四报价，星期一收回，跳过休息日
		end = time.Now().AddDate(0, 0, 4)
		break
	case 5:
		week = "下周一" //星期五报价，星期一收回
		end = time.Now().AddDate(0, 0, 3)
		break
	case 6:
		week = "下周一"
		break

	default:
		week = "下周一"
		break
	}

	body := "ONES内部报价通知："
	endDate := end.Format("2006年01月02日")
	for _, story := range stories {
		numbers = append(numbers, story.Number) //增加编号
		Group1WikiUUID := SmtpService.GetWikiUUID(fmt.Sprintf("#%s ONES交付一组-报价", story.Number))
		Group1WikiLink := fmt.Sprintf("%s/wiki/#/team/RDjYMhKq/space/%s/page/%s",
			global.App.Config.ONES.Base,
			global.App.Config.ONES.SpaceUUID,
			Group1WikiUUID,
		)
		Group2WikiUUID := SmtpService.GetWikiUUID(fmt.Sprintf("#%s ONES交付三组-报价", story.Number))
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
			--------------------------------------------------------------------------------------------------------------
		`,
			story.Number,
			story.Title,
			global.App.Config.ONES.Base,
			story.UUID,
			Group1WikiLink,
			Group2WikiLink,
		) + "\n"
	}
	body += fmt.Sprintf(`【以上需求请在两天内（截止时间：【%s（%s）】）完成，谢谢！】
			==============================================================================================================

			【以下内容，可以前往 https://partner.ones.ai 实例，扭转需求状态为下单后，点击发布按钮后，通知相关服务商！！！】

			ONES需求团队实例列表：
			https://partner.ones.ai/project/#/team/T89vPYaZ/project/UJpAKvZy3ZSeigr1/component/ZIAY1iA9/view/LtxdYvFm   
				
			发布按钮：
			https://partner.ones.ai/project/#/org/SzdYwvt8/setting/business_plugin_service/AppID/YvSxJn0O/InstanceID/4cae55e0/settings-Ry_1
			

			`, endDate, week)

	body += fmt.Sprintf(`
			光恒团队：  
			您好，【%s】新增%d个需求%s号需求，状态已更新为下单中： 
			其中%s号需求请在【%s（%s）】完成报价。 
			需求列表：  
			https://partner.ones.ai/project/#/team/GzPY6Hs8/project/UJpAKvZyVTcsXCle/component/A7s6KQFY/view/2DoUuE6J  

			揽月团队：
			您好，【%s】新增%d个需求%s号需求，状态已更新为下单中：
			其中%s号需求请在【%s（%s）】完成报价。
			需求列表：
			https://partner.ones.ai/project/#/team/P4fGixEe/project/UJpAKvZytcIhwrZ7/component/TRinZhfP/view/MeodL2Ys  

			锦高团队：
			您好，【%s】新增%d个需求%s号需求，状态已更新为下单中：
			其中%s号需求请在【%s（%s）】完成报价。
			需求列表： 
			https://partner.ones.ai/project/#/team/V1FJr5kd/project/UJpAKvZyOo3ZXLng/component/JxMPKM9a/view/Wj7DcqfY  
			
	`,
		startDate, len(numbers), numbers, numbers, endDate, week,
		startDate, len(numbers), numbers, numbers, endDate, week,
		startDate, len(numbers), numbers, numbers, endDate, week,
	)
	body += `

		tips:邮件推荐使用网易灵犀办公打开
	`

	subject := fmt.Sprintf("%s号需求报价", numbers)

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

// GetWikiUUID 得到已询价的UUID
func (s *smtpService) GetWikiUUID(wikiName string) (page string) {
	url := global.App.Config.ONES.Base + "/wiki/api/wiki/team/RDjYMhKq/items/graphql"
	method := "POST"

	searchBody := fmt.Sprintf(
		"{\"query\":\"{ pages(filter:{ name_in:[\\\"%s\\\"],space:{uuid_in:[\\\"%s\\\"], }}){ uuid }}\",\"variables\":{}}",
		wikiName, global.App.Config.ONES.SpaceUUID,
	)

	payload := strings.NewReader(searchBody)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Ones-User-ID", global.App.Config.ONES.UserUUID)
	req.Header.Add("Ones-Auth-Token", global.App.Config.ONES.UserToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	pages := gjson.Get(string(body), "data.pages.#.uuid").Array()
	if len(pages) > 0 {
		return fmt.Sprintf("%s", pages[0])
	}
	return
}

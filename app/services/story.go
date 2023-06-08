package services

import (
	"errors"

	"github.com/VINDA-98/Fasthop/app/common/request"
	"github.com/VINDA-98/Fasthop/app/models"
	"github.com/VINDA-98/Fasthop/global"
)

type storyService struct {
}

var StoryService = new(storyService)

// List 获取需要发布到服务商实例的新实例需求
func (storyService *storyService) List() (stories []models.Story, err error) {
	global.App.DB.Where("is_sync = ?", false).Find(&stories)
	return stories, nil
}

// GetStoryByTitle 通过标题获取用户故事
func (storyService *storyService) GetStoryByTitle(title string) (story models.Story, err error) {
	var result = global.App.DB.Where("title = ?", title).First(&story)
	if result.RowsAffected == 0 {
		err = errors.New("该标题没有对应的用户故事")
	}
	return
}

// GetStoryByNumber 通过标题获取用户故事
func (storyService *storyService) GetStoryByNumber(number string) (story models.Story, err error) {
	var result = global.App.DB.Where("number = ?", number).First(&story)
	if result.RowsAffected == 0 {
		err = errors.New("该需求编号没有对应的用户故事")
	}
	return
}

// PushNewStory 上传新的用户故事
func (storyService *storyService) PushNewStory(params request.Tasks) (stories, badStories []models.Story, err error) {
	var reqStories = params.Stories
	var sendEmails []models.Story
	for i := 0; i < len(reqStories); i++ {
		var story = models.Story{
			UUID:   reqStories[i].UUID,
			Title:  reqStories[i].Title,
			Desc:   reqStories[i].Desc,
			Number: reqStories[i].Number,
			IsSync: false,
		}
		var selectStory = models.Story{}
		var result = global.App.DB.Where("uuid = ?", reqStories[i].UUID).First(&selectStory)
		// 更新发布需求
		if result.RowsAffected != 0 {
			story.IsSync = selectStory.IsSync       // 保证同步状态一致
			if story.Number != selectStory.Number { // 如果编号不一致，说明是新的需求，需要重新同步
				story.IsSync = false
			}
			result = global.App.DB.Where("uuid = ? ", story.UUID).Updates(story)
			if result.RowsAffected == 0 {
				badStories = append(badStories, story)
			} else {
				stories = append(stories, story)
			}
			continue
		}

		//插入新的发布需求
		err = global.App.DB.Create(&story).Error
		if err != nil {
			badStories = append(badStories, story)
		} else {
			stories = append(stories, story)
			sendEmails = append(sendEmails, story) //新插入的需求，需要发送邮件
		}
	}
	//直接发送邮件到相关人员
	SmtpService.SendMail(sendEmails)
	return
}

// UpdateStoryByUUID 通过UUID更新用户故事
func (storyService *storyService) UpdateStoryByUUID(uuid string) (err error) {
	var story = models.Story{
		IsSync: true,
	}
	var result = global.App.DB.Where("uuid = ?", uuid).Updates(&story)
	// 更新发布需求
	if result.RowsAffected == 0 {
		err = errors.New("更新用户故事失败")
	}
	return nil
}

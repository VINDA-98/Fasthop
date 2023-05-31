package app

import (
	"github.com/VINDA-98/Fasthop/app/common/request"
	"github.com/VINDA-98/Fasthop/app/common/response"
	"github.com/VINDA-98/Fasthop/app/services"
	"github.com/gin-gonic/gin"
)

// ListStories 提供服务商实例的查询【用户故事】信息列表接口
func ListStories(c *gin.Context) {
	list, _ := services.StoryService.List()
	response.Success(c, list)
}

// PushStories ONES 发布最新的【用户故事】
func PushStories(c *gin.Context) {
	var form request.Tasks
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if stories, badStories, err := services.StoryService.PushNewStory(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, map[string]interface{}{"stories": stories, "bad_stories": badStories})
	}
}

// UpdateStory 更新用户故事
func UpdateStory(c *gin.Context) {
	err := services.StoryService.UpdateStoryByUUID(c.Param("uuid"))
	if err != nil {
		response.BusinessFail(c, err.Error())
	}
	response.Success(c, "更新用户故事成功")
}

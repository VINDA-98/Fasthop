package request

type Tasks struct {
	Stories []PushStory `form:"stories" json:"stories" binding:"required"`
}

type PushStory struct {
	UUID   string `form:"uuid" json:"uuid" binding:"required"`
	Title  string `form:"title" json:"title" binding:"required"`
	Desc   string `form:"desc" json:"desc" binding:"required"`
	Number string `form:"number" json:"number" binding:"required"`
}

func (PushStory PushStory) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"uuid.required":   "用户故事UUID不能为空",
		"title.required":  "用户故事表头不能为空",
		"desc.required":   "用户故事描述不能为空",
		"number.required": "用户故事编号不能为空",
	}
}

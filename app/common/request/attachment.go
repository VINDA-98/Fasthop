package request

// @Title  request
// @Description  MyGO
// @Author  WeiDa  2023/4/24 14:56
// @Update  WeiDa  2023/4/24 14:56

type AttachmentDownloadForm struct {
	Number   string `json:"number" binding:"required"`
	FileName string `json:"fileName" binding:"required"`
}

func (attachmentDownloadForm AttachmentDownloadForm) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Number.required":   "项目编号不能为空",
		"FileName.required": "附件名称不能为空",
	}
}

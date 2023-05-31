package request

type QueryQuotations struct {
	UUID string `form:"uuid" json:"uuid" binding:"required"`
}

type PushQuotations struct {
	Quotations []PushQuotation `form:"quotations" json:"quotations" binding:"required"`
}

type PushQuotation struct {
	UUID           string  `form:"uuid" json:"uuid" binding:"required"`
	Title          string  `form:"title" json:"title" binding:"required"`
	Desc           string  `form:"desc" json:"desc" binding:"required"`
	PartnerName    string  `form:"partnerName" json:"partnerName" binding:"required"`
	RequirementNo  string  `form:"requirementNo" json:"requirementNo" binding:"required"`
	FileName       string  `form:"fileName" json:"fileName" binding:"required"`
	FilePath       string  `form:"filePath" json:"filePath" binding:"required"`
	Amount         float64 `form:"amount" json:"amount" binding:"required"`
	TotalPersonDay float32 `form:"totalPersonDay" json:"totalPersonDay" binding:"required"`
	IsSync         bool    `form:"is_sync" json:"is_sync" binding:"required"`
}

func (PushQuotation PushQuotation) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"uuid.required":           "用户故事UUID不能为空",
		"title.required":          "用户故事表头不能为空",
		"desc.required":           "用户故事描述不能为空",
		"partnerName.required":    "服务商名称不能为空",
		"requirementNo.required":  "需求编号不能为空",
		"amount.required":         "报价金额不能为空",
		"totalPersonDay.required": "报价人天不能为空",
	}
}

package models

type Quotation struct {
	ID
	UUID           string  `json:"uuid" gorm:"primaryKey;not null;comment:报价需求单的UUID"`
	Title          string  `json:"title" gorm:"not null;comment:报价需求单的标题"`
	Desc           string  `json:"desc" gorm:"not null;comment:报价需求单的描述"`
	PartnerName    string  `json:"partnerName" gorm:"partner_name;not null;comment:报价需求单的服务商名称"`
	RequirementNo  string  `json:"requirementNo" gorm:"requirement_no;not null;comment:报价需求单的需求编号"`
	Amount         float64 `json:"amount" gorm:"amount;not null;comment:报价需求单的报价金额"`
	TotalPersonDay float32 `json:"totalPersonDay" gorm:"total_person_day;not null;comment:报价需求单的报价总人天"`
	RelationUUID   string  `json:"relationUUID" gorm:"relation_uuid;not null;comment:关联的ONES需求用户故事UUID"`
	FileName       string  `json:"fileName" gorm:"filename;not null;comment:Excel报价文件名称"`
	FilePath       string  `json:"filePath" gorm:"filepath;not null;comment:Excel报价文件路径"`
	IsSync         bool    `json:"is_sync" gorm:"is_sync;not null;comment:是否已经同步"`
	Timestamps
	SoftDeletes
}

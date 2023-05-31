package models

import "strconv"

type Story struct {
	ID
	UUID   string `json:"uuid" gorm:"primaryKey;not null;comment:用户故事任务的UUID"`
	Title  string `json:"title" gorm:"not null;comment:用户故事任务的标题"`
	Desc   string `json:"desc" gorm:"size:5000;not null;comment:用户故事任务的描述"`
	Number string `json:"number" gorm:"not null;comment:用户故事任务的编号"`
	IsSync bool   `json:"is_sync" gorm:"is_sync;not null;comment:是否已经同步"`
	Timestamps
	SoftDeletes
}

func (story Story) GetSid() string {
	return strconv.Itoa(int(story.ID.ID))
}

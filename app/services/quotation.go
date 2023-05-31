package services

import (
	"github.com/VINDA-98/Fasthop/app/common/request"
	"github.com/VINDA-98/Fasthop/app/models"
	"github.com/VINDA-98/Fasthop/global"
)

type quotationService struct {
}

var QuotationService = new(quotationService)

// List 获取需要发布到服务商实例的新实例需求
func (quotationService *quotationService) List() (quotation []models.Quotation, err error) {
	err = global.App.DB.Where("is_sync = ?", false).Find(&quotation).Error
	return quotation, err
}

// GetListByRelationUUID 通过用户故事UUID获取对应的报价记录
func (quotationService *quotationService) GetListByRelationUUID(uuid string) (quotation []models.Quotation, err error) {
	err = global.App.DB.Where("relation_uuid = ?", uuid).Find(&quotation).Error
	return quotation, err
}

// GetListByUUID 通过UUID获取对应的报价记录
func (quotationService *quotationService) GetListByUUID(uuid string) (quotation *models.Quotation, err error) {
	err = global.App.DB.Where("uuid = ?", uuid).First(&quotation).Error
	return quotation, err
}

// PushNewQuotation 上传新的报价记录
func (quotationService *quotationService) PushNewQuotation(params request.PushQuotations) (quotations, badQuotations []models.Quotation, err error) {
	var reqQuotations = params.Quotations
	for i := 0; i < len(reqQuotations); i++ {
		var selectQuotation = models.Quotation{}
		var result = global.App.DB.Model(models.Quotation{}).
			Where("uuid = ?", reqQuotations[i].UUID).
			First(&selectQuotation)

		// 更新报价记录
		if result.RowsAffected != 0 {
			// 判断价格或总人天是否发生变化
			if reqQuotations[i].Amount != selectQuotation.Amount || reqQuotations[i].TotalPersonDay != selectQuotation.TotalPersonDay {
				selectQuotation.Amount = reqQuotations[i].Amount
				selectQuotation.TotalPersonDay = reqQuotations[i].TotalPersonDay
				result = global.App.DB.Model(models.Quotation{}).Where("uuid = ? ", selectQuotation.UUID).Updates(selectQuotation)
				if result.RowsAffected == 0 {
					badQuotations = append(badQuotations, selectQuotation)
					continue
				}
			}
			quotations = append(quotations, selectQuotation)
			continue
		}

		// 通过标题获得需求报价单对应的用户故事UUID
		var quotation = models.Quotation{
			UUID:           reqQuotations[i].UUID,
			Title:          reqQuotations[i].Title,
			Desc:           reqQuotations[i].Desc,
			PartnerName:    reqQuotations[i].PartnerName,
			RequirementNo:  reqQuotations[i].RequirementNo,
			Amount:         reqQuotations[i].Amount,
			TotalPersonDay: reqQuotations[i].TotalPersonDay,
			FileName:       reqQuotations[i].FileName,
			FilePath:       reqQuotations[i].FilePath,
		}

		story, err := StoryService.GetStoryByNumber(reqQuotations[i].RequirementNo)
		if err != nil {
			badQuotations = append(badQuotations, quotation)
			continue
		}

		// 设置关联的用户故事UUID
		quotation.RelationUUID = story.UUID

		// 插入报价记录
		err = global.App.DB.Create(&quotation).Error
		if err != nil {
			badQuotations = append(badQuotations, quotation)
		} else {
			quotations = append(quotations, quotation)
		}
	}
	return
}

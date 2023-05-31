package app

import (
	"github.com/VINDA-98/Fasthop/app/common/request"
	"github.com/VINDA-98/Fasthop/app/common/response"
	"github.com/VINDA-98/Fasthop/app/services"
	"github.com/gin-gonic/gin"
)

// ListQuotations  提供ONES实例的查询【报价记录】信息列表接口
func ListQuotations(c *gin.Context) {
	list, _ := services.QuotationService.List()
	response.Success(c, list)
}

// GetQuotationByRelationUUID 提供ONES实例的通过用户故事UUID查询【报价记录】
func GetQuotationByRelationUUID(c *gin.Context) {
	list, _ := services.QuotationService.GetListByRelationUUID(c.Param("relation_uuid"))
	response.Success(c, list)
}

// GetQuotationByUUID 通过UUID查询【报价记录】
func GetQuotationByUUID(c *gin.Context) {
	quotation, _ := services.QuotationService.GetListByUUID(c.Param("uuid"))
	response.Success(c, quotation)
}

// PushQuotation 服务商实例发布最新的【报价记录】
func PushQuotation(c *gin.Context) {
	var form request.PushQuotations
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if quotations, badQuotations, err := services.QuotationService.PushNewQuotation(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, map[string]interface{}{"quotations": quotations, "bad_quotations": badQuotations})
	}
}

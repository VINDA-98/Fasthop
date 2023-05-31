package app

import (
	"fmt"
	"os"

	"github.com/VINDA-98/Fasthop/app/common/response"
	"github.com/VINDA-98/Fasthop/app/services"
	"github.com/gin-gonic/gin"
)

// UploadExcel 上传Excel
func UploadExcel(c *gin.Context) {

	// Multipart form
	file, err := c.FormFile("file")
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	excelName := c.PostForm("excelName")
	number := c.PostForm("number")
	if excelName == "" || number == "" {
		response.ValidateFail(c, "错误的文件参数信息")
		return
	}
	//文件夹路径存储到项目的路径
	fileDir := fmt.Sprintf("static/upload/报价记录/%s", number)

	//文件夹路径拥有执行权限 0777
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	//文件路径
	filePathStr := fmt.Sprintf("static/upload/报价记录/%s/%s", number, excelName)

	// 上传文件至指定目录
	err = c.SaveUploadedFile(file, filePathStr)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, map[string]interface{}{"isUpload": true, "filePath": filePathStr, "errorMsg": ""})
}

// DownloadExcel 下载Excel
func DownloadExcel(c *gin.Context) {
	quotation, _ := services.QuotationService.GetListByUUID(c.Param("attachment_id"))
	if quotation.FilePath == "" {
		response.BusinessFail(c, "对应报价记录的Excel上传路径为空")
		return
	}
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", quotation.UUID+".xlsx"))
	c.File(quotation.FilePath)
}

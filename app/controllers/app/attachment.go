package app

import (
	"fmt"
	"os"

	"github.com/VINDA-98/Fasthop/app/common/request"

	"github.com/VINDA-98/Fasthop/app/common/response"
	"github.com/gin-gonic/gin"
)

var (
	userStoriesDir      = "static/upload/user-stories/attachment/%s"
	userStoriesSavePath = "static/upload/user-stories/attachment/%s/%s"
)

// UploadAttachment 上传attachment
func UploadAttachment(c *gin.Context) {

	// Multipart form
	form, _ := c.MultipartForm()

	files := form.File["upload[]"]

	number := c.PostForm("number")
	if number == "" {
		response.ValidateFail(c, "错误的项目标号信息")
		return
	}

	//文件夹路径存储到项目的路径
	fileDir := fmt.Sprintf(userStoriesDir, number)

	//文件夹路径拥有执行权限 0777
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	successFilePaths := make([]string, 0)
	failedFilePaths := make([]string, 0)
	for _, file := range files {
		//文件路径
		filePathStr := fmt.Sprintf(userStoriesSavePath, number, file.Filename)

		// 上传文件至指定目录
		err = c.SaveUploadedFile(file, filePathStr)
		if err != nil {
			failedFilePaths = append(failedFilePaths, filePathStr)
			continue
		}
		successFilePaths = append(successFilePaths, filePathStr)
	}

	response.Success(
		c, map[string]interface{}{
			"successFilePaths": successFilePaths,
			"failedFilePaths":  failedFilePaths,
			"errorMsg":         "",
		},
	)
}

// DownloadAttachment 下载Attachment  todo 可以优化使用阿里云或者七牛云的下载
func DownloadAttachment(c *gin.Context) {
	var form request.AttachmentDownloadForm
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	path := fmt.Sprintf(userStoriesSavePath, form.Number, form.FileName)
	exist, _ := IsExist(path)
	if !exist {
		response.BusinessFail(c, "文件不存在")
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", form.FileName))
	c.File(path)
}

// AttachmentList 用户故事的所有附件路径列表
func AttachmentList(c *gin.Context) {
	number, _ := c.GetQuery("number")
	if number == "" {
		response.ValidateFail(c, "错误的项目编号")
		return
	}

	// 用户故事目录下所有的文件目录
	fileDir := fmt.Sprintf(userStoriesDir, number)
	// 读取目录下所有的文件
	list, _ := ListDir(fileDir)
	response.Success(
		c, map[string]interface{}{
			"list": list,
		},
	)
}

func ListDir(dirname string) ([]string, error) {
	infos, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(infos))
	for i, info := range infos {
		names[i] = info.Name()
	}
	return names, nil
}

func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

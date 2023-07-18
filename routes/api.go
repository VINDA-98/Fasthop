package routes

import (
	"github.com/VINDA-98/Fasthop/app/controllers/app"
	"github.com/VINDA-98/Fasthop/app/controllers/common"
	"github.com/VINDA-98/Fasthop/app/middleware"
	"github.com/VINDA-98/Fasthop/app/services"
	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	//接口限流
	router.Use(middleware.LimitHandler(10, 1.0))
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}

	userStoryRouter := router.Group("stories").Use(middleware.JWTAuth(services.AppGuardName))
	{
		userStoryRouter.GET("", app.ListStories)
		userStoryRouter.POST("/push", app.PushStories)
		userStoryRouter.POST("/update/:uuid", app.UpdateStory)
	}

	quotationsRouter := router.Group("quotations").Use(middleware.JWTAuth(services.AppGuardName))
	{
		quotationsRouter.GET("", app.ListQuotations)
		quotationsRouter.GET("/:uuid", app.GetQuotationByUUID)
		quotationsRouter.GET("/tasks/:relation_uuid", app.GetQuotationByRelationUUID)
		quotationsRouter.POST("/push", app.PushQuotation)
	}

	// 上传
	uploadFiles := router.Group("upload").Use(middleware.JWTAuth(services.AppGuardName))
	{
		uploadFiles.POST("/excel", app.UploadExcel)
		uploadFiles.POST("/attachment", app.UploadAttachment)
	}

	// 下载
	downloadExcelRouter := router.Group("download").Use(middleware.JWTAuth(services.AppGuardName))
	{
		downloadExcelRouter.POST("/excel/:attachment_id", app.DownloadExcel)
		downloadExcelRouter.GET("/attachments/:number", app.AttachmentList)
		downloadExcelRouter.POST("/attachments", app.DownloadAttachment)
	}

	// 测试接口
	testRouter := router.Group("test").Use(middleware.JWTAuth(services.AppGuardName))
	{
		testRouter.POST("/send", app.SendMailTest)
	}
}

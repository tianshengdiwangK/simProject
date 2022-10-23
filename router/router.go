package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tianshengdiwangK/simProject/basic/log"
	uploader "github.com/tianshengdiwangK/simProject/controller/filecontroller"
	userControl "github.com/tianshengdiwangK/simProject/controller/usercontroller"
	"github.com/tianshengdiwangK/simProject/middleware"
	"go.uber.org/zap"
	//"github.com/axetroy/gin-uploader"
)

var logger *zap.SugaredLogger

func init() {
	logger = log.CwLog()
}

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.GinLogger(log.CwLog()), middleware.GinRecovery(log.CwLog(), true))
	u, _ := uploader.New(router, uploader.TConfig{
		Path:      "upload",
		UrlPrefix: "/api/v1",
		File: uploader.FileConfig{
			Path:      "files",
			MaxSize:   10485760,
			AllowType: []string{"xxx"},
		},
		Image: uploader.ImageConfig{
			Path:    "images",
			MaxSize: 10485760,
			Thumbnail: uploader.ThumbnailConfig{
				Path:      "thumbnail",
				MaxWidth:  300,
				MaxHeight: 300,
			},
		},
	})
	u.Resolve()
	userGroup := router.Group("user")
	{
		//测试mysql
		userGroup.GET("getUser", userControl.GetUser)
		//测试redis
		userGroup.GET("saveUser", userControl.SaveUser)
	}
	return router
}

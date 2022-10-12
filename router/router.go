package main

import (
	"github.com/gin-gonic/gin"
	setup "github.com/tianshengdiwangK/simProject/log"

	"github.com/axetroy/gin-uploader"
)

func init() {
	setup.InitLogger()
}

//func main() {
//	router := gin.New()
//	router.Use(setup.GinLogger(setup.CwLog()), setup.GinRecovery(setup.CwLog(), true))
//	router.GET("/", func(context *gin.Context) {
//		context.JSON(200, "Hello Gin")
//	})
//	err := router.Run(":8008")
//	if err != nil {
//		return
//	}
//}

func main() {
	Router := gin.Default()

	u, err := uploader.New(Router, uploader.TConfig{
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
	if err != nil {
		return
	}

	if err := Router.Run("localhost:9090"); err != nil {
		panic(err)
		return
	}

}

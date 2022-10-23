package usercontroller

import (
	"github.com/gin-gonic/gin"
	setup "github.com/tianshengdiwangK/simProject/basic/log"
	"github.com/tianshengdiwangK/simProject/model"
	"github.com/tianshengdiwangK/simProject/model/reso"
	"github.com/tianshengdiwangK/simProject/service"
	"go.uber.org/zap"
)

var (
	userService *service.UserService
	logger      *zap.SugaredLogger
)

func init() {
	userService = service.NewUserService()
	logger = setup.CwLog()
}

func GetUser(c *gin.Context) {
	username := c.Query("username")

	user, err := userService.QueryByUsername(username)
	if err != nil {
		logger.Error("根据username查询数据库失败")
		c.JSON(500, model.ErrorQueryDatabase(err))
		return
	}
	ret := reso.GetUser{
		ID:       user.ID,
		Username: user.Username,
		Gender:   user.Gender,
		Age:      user.Age,
		Interest: user.Interest,
	}
	c.JSON(200, ret)
}

func SaveUser(c *gin.Context) {
	userId := c.Query("userId")
	err := userService.SaveUserInRedis(userId)
	if err != nil {
		logger.Error("保存到redis失败")
		c.JSON(500, model.ErrorInsertRedis(err))
		return
	}
	c.JSON(200, "success")
}

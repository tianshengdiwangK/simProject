package service

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/tianshengdiwangK/simProject/basic/db"
	redis2 "github.com/tianshengdiwangK/simProject/basic/redis"
	"github.com/tianshengdiwangK/simProject/model/pojo"
	"time"
)

type UserService struct {
	db *gorm.DB
	rd *redis.Client
}

func NewUserService() *UserService {
	dbClient := db.GetDB()
	redisClient := redis2.GetRedis()
	return &UserService{
		db: dbClient,
		rd: redisClient,
	}
}

// QueryByUsername return one user
func (userService *UserService) QueryByUsername(username string) (pojo.User, error) {
	var user = pojo.User{}
	//if err := userService.db.Model(&pojo.User{}).Where("id = ?", 1).Find(&user).Error; err != nil {
	//	return pojo.User{}, err
	//}
	if err := userService.db.First(&user, 1).Error; err != nil {
		return pojo.User{}, err
	}

	return user, nil
}
func (userService *UserService) SaveUserInRedis(userId string) error {
	err := userService.rd.Set("id", userId, time.Minute*30).Err()
	if err != nil {
		return err
	}
	return err
}

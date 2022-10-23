package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	viper2 "github.com/spf13/viper"
	"github.com/tianshengdiwangK/simProject/basic/config"
	"github.com/tianshengdiwangK/simProject/basic/log"
	"github.com/tianshengdiwangK/simProject/model/pojo"
	"go.uber.org/zap"
	"time"
)

var (
	logger *zap.SugaredLogger
	viper  *viper2.Viper
	gormDB *gorm.DB
)

func init() {
	viper = config.NewViper()
	logger = log.CwLog()
	dbType := viper.GetString("database.driver")

	switch dbType {
	case "mysql":
		initMysql()
	default:
		logger.Error("only support mysql")
	}
}

func initMysql() {
	var err error

	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	timeout := viper.GetString("database.timeout")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbname, timeout)
	gormDB, err = gorm.Open(viper.GetString("database.driver"), dsn)

	if err != nil {
		logger.Fatalf("创建数据库连接失败:%v", err)
	}
	gormDB.Value = "xhk"
	//默认不加复数
	gormDB.SingularTable(true)
	//设置连接池
	//空闲
	gormDB.DB().SetMaxIdleConns(20)
	//打开
	gormDB.DB().SetMaxOpenConns(100)
	//超时
	gormDB.DB().SetConnMaxLifetime(time.Second * 30)
	gormDB.AutoMigrate(&pojo.User{})
}
func GetDB() *gorm.DB {
	return gormDB
}

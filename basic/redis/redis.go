package redis

import (
	"github.com/go-redis/redis"
	viper2 "github.com/spf13/viper"
	"github.com/tianshengdiwangK/simProject/basic/config"
	"github.com/tianshengdiwangK/simProject/basic/log"
	"go.uber.org/zap"
	"sync"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
	logger *zap.SugaredLogger
	viper  *viper2.Viper
)

// Init 初始化Redis
func init() {
	m.Lock()
	defer m.Unlock()
	logger = log.CwLog()
	viper = config.NewViper()
	if inited {
		logger.Info("已经初始化过Redis...")
		return
	}

	// 打开才加载
	if viper.GetBool("redis.enabled") {
		logger.Info("初始化Redis...")

		// 加载哨兵模式
		if viper.GetBool("redis.sentinel.enabled") {
			logger.Info("初始化Redis，哨兵模式...")
			initSentinel()
		} else { // 普通模式
			logger.Info("初始化Redis，普通模式...")
			initSingle()
		}

		logger.Info("初始化Redis，检测连接...")

		pong, err := client.Ping().Result()
		if err != nil {
			logger.Fatal(err.Error())
		}
		logger.Info("初始化Redis，检测连接Ping.")
		logger.Info("初始化Redis，检测连接Ping..")
		logger.Infof("初始化Redis，检测连接Ping... %s", pong)
	}
	inited = true
}

// GetRedis 获取redis
func GetRedis() *redis.Client {
	return client
}

func initSentinel() {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    viper.GetString("redis.sentinel.master"),
		SentinelAddrs: viper.GetStringSlice("redis.sentinel.nodes"),
		DB:            viper.GetInt("redis.dbNum"),
		Password:      viper.GetString("redis.password"),
	})

}

func initSingle() {
	client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.conn"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.dbNum"),       // use default DB
	})
}

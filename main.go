package main

import (
	"Pinpin/middleware"
	"Pinpin/routes"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Pinpin/config"
	"Pinpin/global"
	"Pinpin/model"
	"time"
)
var (
	configPrefix = "config_"
	configSuffix = ".json"
	env          *string
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	routes.UsePinpinRouter(r)
	_ = r.Run(":8080")
}

func init() {
	initFlag()
	initConfig()
	initDB()
	initRedis()
}

func initFlag() {
	env = flag.String("env", "prod", "环境")
	flag.Parse()
}

// 读取环境配置文件
func initConfig() {
	configFile := "/etc/pinpin/"+configPrefix + *env + configSuffix
	fmt.Println(configFile)
	global.Config = config.ReadSettingsFromFile(configFile)
}

func initDB() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(10*time.Second))
	go func(ctx context.Context) {
		conf := global.Config
		settings := conf.DbSettings
		connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8",
			settings.Username, settings.Password, settings.Hostname, settings.Dbname)
		var err1 error
		var localDb *gorm.DB
		localDb, err1 = gorm.Open(mysql.Open(connStr), &gorm.Config{})
		if err1 != nil {
			panic("Database connect error," + err1.Error())
		}
		sqlDB, err := localDb.DB()
		if err != nil {
			panic("Database error")
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(10000)
		sqlDB.SetConnMaxLifetime(time.Minute)
		global.DB = localDb
		global.DB.AutoMigrate(&model.Competition_pinpin{},&model.AuthUser{},&model.Follow{},&model.System_notice{},&model.Notice{},model.Report{},&model.ThumbUp{},&model.Reply{})
		cancel()
	}(ctx)

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			fmt.Println("context timeout exceeded")
			panic("Failed to initialize database connection")
		case context.Canceled:
			fmt.Println("context cancelled by force. whole process is complete")
		}
	}
}

func initRedis() {
	redisSettings := global.Config.RedisSettings
	global.RedisClient = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     redisSettings.Address + ":" + redisSettings.Port,
		Password: redisSettings.Password,
		DB:       0,
	})
	_, err := global.RedisClient.Ping().Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("redis连接成功")
	}
}
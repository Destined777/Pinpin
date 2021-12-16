package global

import (
	"Pinpin/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)
var (
	Config      *config.Config
	DB          *gorm.DB
	RedisClient *redis.Client
)
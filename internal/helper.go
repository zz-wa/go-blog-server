package blog_r

import (
	"blog_r/internal/global"
	"blog_r/internal/model"
	"context"
	"os"
	"regexp"

	"github.com/glebarez/sqlite"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitLogger() {
	var level zapcore.Level
	switch global.Conf.Log.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var encoder zapcore.Encoder
	if global.Conf.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)

}

var dsnPasswordRe = regexp.MustCompile(`password=[^\s]+`)

func maskDsn(dsn string) string {
	if dsn == "" {
		return dsn
	}
	return dsnPasswordRe.ReplaceAllString(dsn, "password=******")
}

func InitDatabase(conf *global.Config) *gorm.DB {
	dbType := conf.DbType()
	dsn := conf.DbDSN()

	var db *gorm.DB
	var err error

	var level logger.LogLevel
	switch conf.Server.DbLogMode {
	case "silent":
		level = logger.Silent
	case "info":
		level = logger.Info
	case "warn":
		level = logger.Warn
	case "error":
		level = logger.Error
	default:
		level = logger.Info
	}

	config := &gorm.Config{
		Logger:                                   logger.Default.LogMode(level),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	switch dbType {
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), config)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), config)
	default:
		zap.L().Error("不支持的数据库类型", zap.String("dbType", dbType))
	}

	if err != nil {
		zap.L().Error("连接数据库失败", zap.String("dbType", dbType), zap.Error(err))
		return nil
	}
	global.DB = db
	zap.L().Info(
		"成功连接数据库",
		zap.String("type", dbType),
		zap.String("dsn", maskDsn(dsn)),
	)

	if conf.Server.DbAutoMigrate {
		if err := model.MakeMigration(db); err != nil {
			zap.L().Error("自动迁移失败", zap.Error(err))
			return db
		}
		zap.L().Info("自动迁移成功")
	}
	return db
}

func InitRedis(conf *global.Config) error {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	global.Redis = client
	zap.L().Info("Redis 连接成功", zap.String("pong", pong), zap.String("addr", conf.Redis.Addr))
	return nil
}

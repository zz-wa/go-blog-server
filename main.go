package main

import (
	blog_r "blog_r/internal"
	"blog_r/internal/global"
	"blog_r/internal/router"
	"blog_r/internal/service/user"
	"flag"

	"go.uber.org/zap"
)

func main() {
	configPath := flag.String("config", "./internal/global/config.yaml", "Path to the configuration file")

	flag.Parse()
	conf := global.ReadConf(*configPath)
	blog_r.InitLogger()
	db := blog_r.InitDatabase(conf)
	if db == nil {
		zap.L().Fatal("数据库初始化失败")
	}
	if err := user.InitAdmin(conf); err != nil {
		zap.L().Fatal("管理员初始化失败", zap.Error(err))
	}
	err := blog_r.InitEnforce()
	if err != nil {
		zap.L().Fatal("Casbin初始化失败")
	}
	enforce, err := global.Enforcer.Enforce("admin", "/api/v1/admin/categories", "POST")
	if err != nil {
		zap.L().Fatal("casbin 验证出错", zap.Error(err))
	}
	if !enforce {
		zap.L().Fatal("验证失败")
	}
	enforce, err = global.Enforcer.Enforce("reader", "/api/v1/admin/categories", "POST")
	if err != nil {
		zap.L().Fatal("casbin 验证出错", zap.Error(err))
	}
	if enforce {
		zap.L().Fatal("验证失败")
	}

	err = blog_r.InitRedis(conf)
	if err != nil {
		zap.L().Fatal("redis connection error")
	}
	zap.L().Info("程序运行成功", zap.String("configPath", *configPath))
	port := conf.Server.Port
	if port == "" {
		port = "8080"
	}

	zap.L().Info("服务器正在运行", zap.String("port", port))
	e := router.NewRouter()
	if err := e.Start(":" + port); err != nil {
		zap.L().Error("服务器启动失败", zap.Error(err))
	}
}

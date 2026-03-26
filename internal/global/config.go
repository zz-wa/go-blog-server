package global

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct { //定义了整个应用程序所需要的配置项，包括服务器配置、日志配置、JWT配置、数据库配置、Redis配置、Session配置、Email配置、验证码配置和文件上传配置等。
	Server struct {
		Model         string
		Port          string
		DbType        string
		DbAutoMigrate bool
		DbLogMode     string
	}
	Log struct {
		Level     string
		Prefix    string
		Format    string
		Directory string
	}
	JWT struct {
		Secret string //密钥
		Expire int64  //过期时间
		Issuer string //发行者
	}
	Pgsql struct {
		Host     string
		Port     string
		User     string
		Config   string //额外的连接参数
		Password string
		DbName   string
	}
	Sqlite struct {
		Dsn string //数据源名称
	}
	Redis struct {
		DB       int
		Addr     string //地址
		Password string //密码
	}
	Session struct {
		Name   string
		Salt   string //盐值
		MaxAge int    //过期时间，单位为秒
	}
	Email struct {
		Form     string
		Host     string
		Port     int
		SmtpPass string
		SmtpUser string
	}
	Captcha struct {
		SendEmail  bool
		ExpireTime int
	}
	Upload struct {
		Size      int
		OssType   string //上传类型，支持本地和七牛云
		Path      string
		StorePath string
	}
	Qiniu struct {
		ImgPath       string
		Zone          string
		Bucket        string
		AccessKey     string
		SecretKey     string
		UseHTTPS      bool
		UseCdnDomains bool
	}
	Admin struct {
		Username string
		Email    string
		Password string
	}
}

var Conf *Config

func GetConfig() *Config {
	if Conf == nil {
		log.Panic("配置文件未进行初始化")
		return nil
	}
	return Conf
}

func ReadConf(path string) *Config {

	v := viper.New()
	v.SetConfigFile(path) //设置配置文件路径
	v.AutomaticEnv()      //允许从环境变量中读取配置项
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		panic("Failed to read the configuration file" + err.Error())

	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic("unmarshall error" + err.Error())
	}
	log.Println("success to read the configuration file ")

	return Conf
}

func (*Config) DbType() string {
	if Conf.Server.DbType == "" {
		Conf.Server.DbType = "sqlite" //默认使用sqlite数据库
	}
	return Conf.Server.DbType
}

func (*Config) DbDSN() string {
	switch Conf.Server.DbType {
	case "postgres":
		conf := Conf.Pgsql //获取PostgreSQL数据库的连接信息
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s", conf.Host, conf.User, conf.Password, conf.DbName, conf.Port, conf.Config)
		//
		//host=127.0.0.1 user=root password=qa.86.... dbname=gva_blog port=5432 sslmode=disable
	case "sqlite":
		return Conf.Sqlite.Dsn
	default:
		Conf.Server.DbType = "sqlite"
		if Conf.Sqlite.Dsn == "" {
			Conf.Sqlite.Dsn = "file::memory"
		}
		return Conf.Sqlite.Dsn
	}
}

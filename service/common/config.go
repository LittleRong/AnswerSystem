package common

import (
	"os"
	"strings"
	"time"
	"context"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

type Config struct {
	Name string
}

var (
	cfg = pflag.StringP("config", "c", "", "service's config")
)

func Init(serviceName string) (micro.Service,error) {
	//配置初始化
	pflag.Parse()
	c := Config {
		Name: *cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return nil,err
	}

	c.initDatabase()//初始化数据库
	service := c.initServiceRegistry(serviceName)//初始化consul
	c.initLog()//初始化日志


	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return service,nil
}


//读取配置文件内容
func (this *Config) initConfig() error {

	if this.Name != "" {
		//若配置文件不存在，抛出报错
		if _, err := os.Stat(this.Name); os.IsNotExist(err) {
			panic(err)
		}

		viper.SetConfigFile(this.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("../../conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	viper.AutomaticEnv()        // 读取匹配的环境变量
	viper.SetEnvPrefix("WEB")   // 读取环境变量的前缀为WEB
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	//监控及热加载
	this.watchConfig()

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed: %s")
	})
}

func (this *Config)initDatabase() error {

	//用户名:密码@数据库地址+名称?字符集,root:password123@tcp(localhost:3306)/problem?charset=utf8
	dataSource := viper.GetString("database.user") + ":" + viper.GetString("database.pwd") + "@" + viper.GetString("database.protol") + "(" + viper.GetString("database.host") + ")" + "/" + viper.GetString("database.name") + "?charset=" + viper.GetString("database.charset")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = viper.GetBool("database.debug")
	// 自动建表
	orm.RunSyncdb("default", false, true)

	return nil
}

func (this *Config)initServiceRegistry(serviceName string) micro.Service{
	//create service
	service := micro.NewService(micro.Name(serviceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		micro.WrapHandler(logWrapper),
		micro.Registry(consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			viper.GetString("consul.host")+":"+viper.GetString("consul.port"),
		}
	})))

	//init
	service.Init()

	return service
}

func (this *Config) initLog() {
	passLagerCfg := log.PassLagerCfg {
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

// 实现server.HandlerWrapper接口
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Debug("server request", lager.Data{
			req.Service():req.Endpoint(),
		})
		return fn(ctx, req, rsp)
	}
}
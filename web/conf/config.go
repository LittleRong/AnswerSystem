package conf

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error{
	c := Config{Name:cfg}

	//读取配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	//监控及热加载
	c.Watch()

	return nil
}

//读取配置文件内容
func (this *Config) initConfig() error{
	if this.Name != "" {
		//若配置文件不存在，抛出报错
		if _,err:=os.Stat(this.Name);os.IsNotExist(err){
			panic(err)
		}

		viper.SetConfigFile(this.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("../conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	viper.AutomaticEnv() // 读取匹配的环境变量
	viper.SetEnvPrefix("WEB") // 读取环境变量的前缀为WEB
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}


//监控及热加载
func (this *Config) Watch(){
	viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	log.Printf("Config file changed: %s", e.Name)
	//})
}
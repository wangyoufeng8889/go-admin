package config
import "github.com/spf13/viper"

type Aliyunconfig struct {
	Uid string
	AccessKey string
	AccessSecret string
	Region string
	ClientId string
}

func InitAliyunconfig(cfg *viper.Viper) *Aliyunconfig {
	return &Aliyunconfig{
		Uid:          cfg.GetString("uid"),
		AccessKey:          cfg.GetString("accessKey"),
		AccessSecret:          cfg.GetString("accessSecret"),
		Region:          cfg.GetString("region"),
		ClientId:          cfg.GetString("clientId"),
	}
}

var AliyunConfig = new(Aliyunconfig)

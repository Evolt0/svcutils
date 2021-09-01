package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"reflect"
)

func triggerInitials(rv reflect.Value) {
	rv = reflect.Indirect(rv)

	if rv.Kind() == reflect.Struct {
		for i := 0; i < rv.NumField(); i++ {
			value := rv.Field(i)

			if conf, ok := value.Interface().(interface{ Init() }); ok {
				conf.Init()
			}
		}
	} else {
		if conf, ok := rv.Interface().(interface{ Init() }); ok {
			conf.Init()
		}
	}
}

func ConfP(data interface{}) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	logrus.Println(wd)
	v := viper.New()
	v.SetConfigName("local")
	v.AddConfigPath(wd + "/config") //文件所在的目录路径
	v.SetConfigType("yml")          //这里是文件格式类型
	err = v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	configs := v.AllSettings()
	for k, val := range configs {
		v.SetDefault(k, val)
	}
	logrus.Println(configs)
	err = v.Unmarshal(data) //反序列化至结构体
	if err != nil {
		panic(err)
	}
	logrus.Println(data)
	rv := reflect.ValueOf(data)
	triggerInitials(rv)
}

package main

import (
	"time"

	"github.com/lunny/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var runtime_conf struct {
	DB string `json:"db"`
}

func main() {
	// alternatively, you can create a new viper instance.
	var runtime_viper = viper.New()

	runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
	runtime_viper.SetConfigType("yaml") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop"

	// read from remote config the first time.
	err := runtime_viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}

	// unmarshal config
	runtime_viper.Unmarshal(&runtime_conf)

	// open a goroutine to watch remote changes forever
	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			// currently, only tested with etcd support
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				log.Errorf("unable to read remote config: %v", err)
				continue
			}

			// unmarshal new config into our runtime config struct. you can also use channel
			// to implement a signal to notify the system of the changes
			runtime_viper.Unmarshal(&runtime_conf)
		}
	}()
	time.Sleep(time.Hour)
}

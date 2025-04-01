package config

import (
	"fmt"

	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// 配置结构定义
type TestSt struct {
	Name string `json:"name"`
}

type Config struct {
	rest.RestConf
	UserRpcConf zrpc.RpcClientConf
	Auth        struct {
		AccessSecret string
		AccessExpire int
	}
	Etcd struct {
		Host string
	}
}

func (c *Config) Sub() {
	ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
		Hosts: []string{c.Etcd.Host}, // etcd 地址
		Key:   "test",                // 配置key
	})
	cc := configurator.MustNewConfigCenter[TestSt](configurator.Config{
		Type: "json", // 配置值类型：json,yaml,toml
	}, ss)
	// 获取配置
	// 注意: 配置如果发生变更，调用的结果永远获取到最新的配置
	cfg, err := cc.GetConfig()
	if err != nil {
		fmt.Println(err)
	}
	// 打印配置
	fmt.Println(cfg)
}

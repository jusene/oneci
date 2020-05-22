package utils

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"log"
	"zjhw.com/oneci/config"
)

func consulClient(config *config.ConsulConfig) (*consulApi.Client, error) {
	conf := consulApi.DefaultConfig()
	conf.Address = fmt.Sprintf("%s:%d", config.Address, config.Port)
	client, err := consulApi.NewClient(conf)
	return client, err
}

func GetKV(config *config.ConsulConfig, key string) (*consulApi.KVPair, error) {
	client, err := consulClient(config)
	if err != nil {
		log.Fatalf("**** 创建consul client失败: %v", err)
	}
	KVPair, _, err := client.KV().Get(key, &consulApi.QueryOptions{})
	return KVPair, err
}
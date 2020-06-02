package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"zjhw.com/oneci/config"
)

func DeployKube(conf *config.ConsulConfig, app, version, project, env, ty, arch string, timestamp int64) {
	if ty == "nil" {
		ty = ""
	}

	log.Printf("**** 获取 %s 项目的基本配置", project)
	c := func() *config.AppConfig {
		value, err := GetKV(conf, fmt.Sprintf("/oneci/config/%s", project))
		if err != nil {
			log.Fatalf("**** 获取配置失败, key: %s: %v", fmt.Sprintf("/oneci/config/%s", project), err)
		}
		kubeConf := value.Value
		//fmt.Printf("*** 调试信息\n%s", string(consulConf))
		c := new(config.AppConfig)
		err = yaml.Unmarshal(kubeConf, c)
		if err != nil {
			panic(err)
		}
		return c
	}()

	singleAppConfig := CheckOutAppConfig(app, c)
	singleAppConfig.VERSION = version
	singleAppConfig.PROJECT = project
	singleAppConfig.ARCH = arch
	singleAppConfig.TYPE = ty
	singleAppConfig.ENV = env
	singleAppConfig.TIMESTAMP = timestamp

	if singleAppConfig.Debug {
		switch env {
		case "dev":
			singleAppConfig.DEBUGPORT = singleAppConfig.PORT + 10000
		case "test":
			singleAppConfig.DEBUGPORT = singleAppConfig.PORT + 20000
		case "pre":
			singleAppConfig.DEBUGPORT = singleAppConfig.PORT + 30000
		case "prod":
			singleAppConfig.DEBUGPORT = singleAppConfig.PORT + 40000
		default:
			log.Fatalf("unknown env")
		}
	}

	log.Printf("**** 获取 %s 项目 %s 应用的配置", project, app)
	value, err := GetKV(conf, fmt.Sprintf("/oneci/template/%s/%s", project, app))
	if err != nil {
		log.Fatalf("**** 获取配置失败, key: %s: %v", fmt.Sprintf("/oneci/config/%s/%s", project, app), err)
	}

	os.MkdirAll("kube", 0755)
	target, err := os.OpenFile("kube/"+fmt.Sprintf("%s-%s%s.yml", app, env, ty), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("文件创建失败 %s", "kube/"+fmt.Sprintf("%s-%s%s.yml", app, env, ty))
	}
	defer target.Close()
	Render(string(value.Value), target, singleAppConfig)

	SFTPut(config.HostConf.Host, config.HostConf.Username, config.HostConf.Password, config.HostConf.Port,
		"kube/"+fmt.Sprintf("%s-%s%s.yml", app, env, ty),
		fmt.Sprintf("/ddhome/k8s/%s/%s-%s%s.yml", project, app, env, ty))

	SSHExec(config.HostConf.Host, config.HostConf.Username, config.HostConf.Password, config.HostConf.Port,
		"/usr/local/bin/kubectl apply -f "+fmt.Sprintf("/ddhome/k8s/%s/%s-%s%s.yml", project, app, env, ty))
}

package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"zjhw.com/oneci/config"
)

// 根据consul的应用配置检出相对应的应用配置
func CheckOutAppConfig(app string, conf *config.AppConfig) *config.AppInfo {
	for _, c := range conf.Apps {
		if c.APP == app {
			return &c
		}
	}
	return nil
}

// 准备编译docker的准备工作
func PreJavaDocker(app, project, version, env, arch, ty, dockerfile, entrypoint string) {
	workerSpace := strings.Join([]string{app, "docker"}, "_")
	// 创建打包使用的工作目录
	if err := os.MkdirAll(workerSpace, 0755); err != nil {
		log.Fatalf("**** 创建目录失败 %s; %v", strings.Join([]string{app, "docker"}, "_"), err)
	}

	// 查找编译出的jar包
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && strings.Contains(path, fmt.Sprintf("%s/target", app)) {
			if file, err := filepath.Glob(fmt.Sprintf("%s/*.jar", path)); err == nil && len(file) != 0 {
				log.Printf("**** 应用 %s 找到相应的jar包: %s", app, file[0])
				jarSplit := strings.Split(file[0], "-")
				jarVersion := strings.TrimRight(jarSplit[len(jarSplit)-1], ".jar")
				destFile, _ := os.Create(strings.Join([]string{workerSpace, filepath.Base(file[0])}, "/"))
				defer destFile.Close()
				srcFile, _ := os.Open(file[0])
				defer srcFile.Close()
				io.Copy(destFile, srcFile)

				// 获取项目应用的配置
				log.Printf("**** 获取 %s 项目的基本配置", project)
				conf := func() *config.AppConfig {
					value, err := GetKV(config.Conf, fmt.Sprintf("/oneci/config/%s", project))
					if err != nil {
						log.Fatalf("**** 获取配置失败, key: %s: %v", fmt.Sprintf("/oneci/config/%s", project), err)
					}
					consulConf := value.Value
					//fmt.Printf("*** 调试信息\n%s", string(consulConf))
					c := new(config.AppConfig)
					err = yaml.Unmarshal(consulConf, c)
					if err != nil {
						panic(err)
					}
					return c
				}()
				singleAppConfig := CheckOutAppConfig(app, conf)
				//fmt.Println(singleAppConfig)
				singleAppConfig.VERSION = version
				singleAppConfig.PROJECT = project
				singleAppConfig.JARVERSION = jarVersion
				singleAppConfig.ARCH = arch

				// 特殊需求 部署方式
				if ty != "nil" {
					singleAppConfig.TYPE = ty
				}

				// 根据部署环境查找websocket port
				if len(singleAppConfig.WSPort) != 0 {
					for _, wsinfo := range singleAppConfig.WSPort {
						if wsinfo.ENV == env {
							singleAppConfig.WS = wsinfo.PORT
						}
					}
				}

				// 根据配置中心是否生成debug接口，规则dev: port+10000, test: port+20000, pre: port+30000, prod: port+40000
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

				// 生成应用dockerfile
				log.Printf("**** 应用 %s 生成相应的Dockerfile", app)
				value, err := GetKV(config.Conf, dockerfile)
				if err != nil {
					log.Fatalf("**** 应用 %s 获取模板失败, key: %s: %v", app, dockerfile, err)
				}

				target, err := os.OpenFile(workerSpace+"/Dockerfile", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatalf("文件创建失败 %s", workerSpace+"/Dockerfile")
				}
				defer target.Close()
				Render(string(value.Value), target, singleAppConfig)

				// 生成应用entrypoint脚本
				log.Printf("**** 应用 %s 生成相应的docker-entrypoint.sh", app)
				value, err = GetKV(config.Conf, entrypoint)
				if err != nil {
					log.Fatalf("**** 应用 %s 获取模板失败, key: %s: %v", app, entrypoint, err)
				}
				entry, err := os.OpenFile(workerSpace+"/docker-entrypoint.sh", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatalf("文件创建失败 %s", workerSpace+"/docker-entrypoint.sh")
				}
				defer entry.Close()
				Render(string(value.Value), entry, singleAppConfig)

				// 根据配置中心复制脚本打包进容器
				if singleAppConfig.BIN {
					filepath.Walk(path+"/bin", func(path string, info os.FileInfo, err error) error {
						if !info.IsDir() {
							os.MkdirAll(strings.Join([]string{workerSpace, "bin"}, "/"), 0755)
							destFile, _ := os.Create(strings.Join([]string{workerSpace, "bin", filepath.Base(path)}, "/"))
							defer destFile.Close()
							srcFile, _ := os.Open(path)
							defer srcFile.Close()
							io.Copy(destFile, srcFile)
						}
						return nil
					})
				}

				// 根据配置中心拉取字符库进容器
				if singleAppConfig.FONT {
					filepath.Walk("/ddhome/fonts", func(path string, info os.FileInfo, err error) error {
						os.MkdirAll(strings.Join([]string{workerSpace, "fonts"}, "/"), 0755)
						destFile, _ := os.Create(strings.Join([]string{workerSpace, "fonts", filepath.Base(path)}, "/"))
						defer destFile.Close()
						srcFile, _ := os.Open(path)
						defer srcFile.Close()
						io.Copy(destFile, srcFile)
						return nil
					})
				}

				// 根据配置中心拉取证书进容器
				if singleAppConfig.CERT {
					filepath.Walk("/ddhome/cert", func(path string, info os.FileInfo, err error) error {
						os.MkdirAll(strings.Join([]string{workerSpace, "cert"}, "/"), 0755)
						destFile, _ := os.Create(strings.Join([]string{workerSpace, "cert", filepath.Base(path)}, "/"))
						defer destFile.Close()
						srcFile, _ := os.Open(path)
						defer srcFile.Close()
						io.Copy(destFile, srcFile)
						return nil
					})
				}

				// 根据配置中心将resource拉取进容器
				if singleAppConfig.RESOURCE {
					filepath.Walk(path+"/resources", func(path string, info os.FileInfo, err error) error {
						if !info.IsDir() {
							os.MkdirAll(strings.Join([]string{workerSpace, "resources"}, "/"), 0755)
							destFile, _ := os.Create(strings.Join([]string{workerSpace, "resources", filepath.Base(path)}, "/"))
							defer destFile.Close()
							srcFile, _ := os.Open(path)
							defer srcFile.Close()
							io.Copy(destFile, srcFile)
						}
						return nil
					})
				}
				return nil
			}
		}
		return err
	})
}

func PreJavaScriptDocker() {

}

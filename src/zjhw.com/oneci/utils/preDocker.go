package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"zjhw.com/oneci/config"
)

func PreJavaDocker(app, project, version, dockerfile, entrypoint string) {
	workerSpace := strings.Join([]string{app, "docker"}, "_")
	// 创建打包使用的工作目录
	if err := os.MkdirAll(workerSpace, 0755); err != nil {
		log.Fatalf("**** 创建目录失败 %s; %v", strings.Join([]string{app, "docker"}, "_"), err)
	}

	// 查找编译出的jar包
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && strings.Contains(path, fmt.Sprintf("%s\\target", app)) {
			if file, err := filepath.Glob(fmt.Sprintf("%s\\*.jar", path)); err == nil {
				log.Printf("**** 应用 %s 找到相应的jar包: %s", app, file[0])
				jarSplit := strings.Split(file[0], "-")
				jarVersion := strings.TrimRight(jarSplit[len(jarSplit)-1], ".jar")
				fmt.Println(jarVersion)
				destFile, _ := os.Create(strings.Join([]string{workerSpace, filepath.Base(file[0])}, "/"))
				defer destFile.Close()
				srcFile, _ := os.Open(file[0])
				defer srcFile.Close()
				io.Copy(destFile, srcFile)

				log.Printf("**** 应用 %s 生成相应的Dockerfile", app)
				value, err := GetKV(config.Conf, dockerfile)
				if err != nil {
					log.Fatalf("**** 应用 %s 获取模板失败, key: %s: %v", app, dockerfile, err)
				}
				fmt.Println(string(value.Value))

				log.Printf("**** 应用 %s 生成相应的docker-entrypoint.sh", app)
				value, err = GetKV(config.Conf, entrypoint)
				if err != nil {
					log.Fatalf("**** 应用 %s 获取模板失败, key: %s: %v", app, entrypoint, err)
				}
				fmt.Println(string(value.Value))
				return nil
			}
		}
		return err
	})
}


func PreJavaScriptDocker() {

}
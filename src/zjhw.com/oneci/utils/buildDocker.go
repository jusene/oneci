package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func BuildDocker(app, version, project, env, ty string, timestamp int64) {
	if ty == "nil" {
		ty = ""
	}

	currentPath, _ := filepath.Abs(".")
	os.Chdir(strings.Join([]string{app, "docker"}, "_"))
	cmd := exec.Command("docker", "build", "-t", fmt.Sprintf("harbor.zjhw.com/%s-%s/%s:%s.%d%s",
		project, env, app, version, timestamp, ty), ".")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()

	// 创建一个流来读取管道内的内容，一行一行读
	reader := bufio.NewReader(stdout)

	for {
		// 以换行符作为一行结尾
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		log.Print(line)
	}
	cmd.Wait()

	os.Chdir(currentPath)
}

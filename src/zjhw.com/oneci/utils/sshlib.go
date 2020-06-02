package utils

import (
	"bytes"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func sshConnect(host, username, password string, port int) *ssh.Client {
	conf := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	addr := fmt.Sprintf("%s:%d", host, port)

	sshClient, err := ssh.Dial("tcp", addr, conf)
	if err != nil {
		log.Fatal("创建ssh client失败", err)
	}
	//defer sshClient.Close()
	return sshClient
}

func SSHExec(host, username, password string, port int, command string) {
	sshClient := sshConnect(host, username, password, port)
	defer sshClient.Close()
	// 创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session失败", err)
	}
	defer session.Close()

	var stdOut, stdErr bytes.Buffer
	session.Stdout = &stdOut
	session.Stderr = &stdErr

	session.Run(command)
	if stdErr.String() != "" {
		log.Fatal("err: ", stdErr.String())
	}
	log.Println(stdOut.String())
}

func SFTPut(host, username, password string, port int, src, dest string) {
	sshClient := sshConnect(host, username, password, port)
	defer sshClient.Close()

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatal("创建sftp client失败", err)
	}
	defer sftpClient.Close()

	// src 文件
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal("打开源文件错误:", err)
	}
	defer srcFile.Close()

	// dest 文件
	destFile, err := sftpClient.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	destFile.Write(ff)
	log.Printf("%s => %s", src, dest)
}

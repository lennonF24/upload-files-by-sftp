package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {

	host := "192.168.70.132"
	port := 22
	user := "gouserftp"
	password := ""

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		fmt.Println("Failed to connect to SSH server:", err)
		return
	}
	defer conn.Close()

	// Open SFTP session
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println("Failed to open SFTP session:", err)
		return
	}
	defer sftpClient.Close()

	localFile, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("Failed to open local file:", err)
		return
	}
	defer localFile.Close()

	remoteFile, err := sftpClient.Create("/data/sftp/file.txt")
	if err != nil {
		fmt.Println("Failed to create remote file:", err)
		return
	}
	defer remoteFile.Close()

	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		fmt.Println("Failed to upload file:", err)
		return
	}
}

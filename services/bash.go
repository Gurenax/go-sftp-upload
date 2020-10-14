package services

import (
	"fmt"
	"log"
	"os/exec"

	m "go-sftp-upload/models"
)

// Scp - Run the Scp command using Bash and sshpass
func Scp(env m.Credentials) {
	// Start bash
	bashCommand := exec.Command("bash")
	writer, _ := bashCommand.StdinPipe()
	bashCommand.Start()

	// Execute scp with sshpass
	cmdString := fmt.Sprintf("sshpass -p %s scp -r %s %s@%s:%s", env.Password, env.Src, env.Username, env.Server, env.Dest)
	// fmt.Println(cmdString)
	writer.Write([]byte(cmdString + "\n"))

	// Exit from bash
	writer.Write([]byte("exit" + "\n"))

	// Output bash errors if any
	bashError := bashCommand.Wait()
	if bashError != nil {
		log.Fatal(bashError)
	}
}

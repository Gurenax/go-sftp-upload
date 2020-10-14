package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func checkSSHPass() {
	cmd := exec.Command("which", "sshpass")
	_, err := cmd.CombinedOutput()
	if err != nil {
		// log.Fatal(err)
		fmt.Println("You need to install sshpass to continue")
		fmt.Println("e.g. brew install esolitos/ipa/sshpass (Mac OS)")
		fmt.Println("e.g. apt-get install sshpass (Ubuntu)")
		os.Exit(1)
	}
}

func loadEnv() (string, string, string, string, string) {
	// Load credentials from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	password := os.Getenv("PASSWORD")
	src := fmt.Sprintf("\"%s\"", os.Getenv("SRC"))
	username := os.Getenv("USERNAME")
	server := os.Getenv("SERVER")
	dest := os.Getenv("DEST")

	return password, src, username, server, dest
}

func main() {
	// Check if sshpass exists
	checkSSHPass()

	// Get credentials from .env file
	password, src, username, server, dest := loadEnv()

	// Start bash
	bashCommand := exec.Command("bash")
	writer, _ := bashCommand.StdinPipe()
	bashCommand.Start()

	// Execute scp with sshpass
	cmdString := fmt.Sprintf("sshpass -p %s scp -r %s %s@%s:%s", password, src, username, server, dest)
	fmt.Println(cmdString)
	writer.Write([]byte(cmdString + "\n"))

	// Exit from bash
	writer.Write([]byte("exit" + "\n"))

	// Output bash errors if any
	bashError := bashCommand.Wait()
	if bashError != nil {
		log.Fatal(bashError)
	}
}

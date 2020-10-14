package main

import (
	h "go-sftp-upload/helpers"
	s "go-sftp-upload/services"
)

func main() {
	// Check if sshpass binary exists
	h.CheckBinaryExists("sshpass")

	// Get credentials from .env file
	env := h.LoadEnv()

	// Execute scp with bash service
	s.Scp(env)
}

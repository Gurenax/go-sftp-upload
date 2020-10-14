package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckBinaryExists(binaryName string) {
	cmd := exec.Command("which", binaryName)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("You need to install %s to continue\n", binaryName)
		os.Exit(1)
	}
}

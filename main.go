package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	apps := []string{"Google Chrome", "Visual Studio Code", "Postman", "MongoDB Compass", "WhatsApp"}

	for _, app := range apps {
		err := openApplication(app)
		if err != nil {
			fmt.Printf("Error opening %s: %v\n", app, err)
		} else {
			fmt.Printf("Opened %s successfully\n", app)
		}
	}
}

func openApplication(appName string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", "-a", appName)
	case "windows":
		cmd = exec.Command("start", appName)
	case "linux":
		cmd = exec.Command(appName)
	}
	return cmd.Start()
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", os.Args[1])

	default:
		return
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

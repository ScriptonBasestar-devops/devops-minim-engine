package util

import "os/exec"

func RunCommands(commands []string) {
	for _, command := range commands {
		exec.Command(command)
	}
}

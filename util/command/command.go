package command

import (
	"os/exec"
)

func RunCmd(cmd string) (string, error) {
	cmd1 := exec.Command("/bin/bash", "-c", cmd)
	output, err := cmd1.CombinedOutput()
	return string(output), err
}

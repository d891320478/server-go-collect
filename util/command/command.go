package command

import "os/exec"

func RunCmd(cmd string) (string, error) {
	cmd1 := exec.Command("/bin/bash", "-c", cmd)
	out1, err := cmd1.Output()
	return string(out1), err
}

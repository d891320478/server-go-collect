package command

import (
	"bytes"
	"os/exec"
)

func RunCmd(cmd string) (string, error) {
	buf := bytes.NewBuffer(nil)
	cmd1 := exec.Command("/bin/bash", "-c", cmd)
	cmd1.Stdout = buf
	cmd1.Stderr = buf
	err := cmd1.Run()
	if err != nil {
		return "", err
	}
	return buf.String(), err
}

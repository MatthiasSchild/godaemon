//go:build linux || darwin

package godaemon

import (
	"os"
	"os/exec"
	"syscall"
)

func startProcess(envVariable string) error {
	fileName := os.Args[0]
	env := envVariable + "=1"

	cmd := exec.Command(fileName)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	cmd.Env = append(os.Environ(), env)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func stopProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Signal(syscall.SIGTERM)
}

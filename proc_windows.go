//go:build windows

package godaemon

import (
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

func startProcess(envVariable string) error {
	fileName := os.Args[0]
	env := envVariable + "=1"

	cmd := exec.Command(fileName)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	cmd.Env = append(os.Environ(), env)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func stopProcess(pid int) error {
	handle, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		return err
	}
	defer windows.CloseHandle(handle)

	return windows.TerminateProcess(handle, uint32(syscall.SIGTERM))
}

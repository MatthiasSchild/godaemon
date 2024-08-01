package godaemon

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func writePidFile(serviceName string) error {
	pid := os.Getpid()
	fileName := serviceName + ".pid"

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%d\n", pid)
	if err != nil {
		return err
	}

	return nil
}

func readPidFile(serviceName string) (int, error) {
	fileName := serviceName + ".pid"

	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	pidString := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(pidString)
	if err != nil {
		return 0, err
	}

	return pid, nil
}

func deletePidFile(serviceName string) error {
	fileName := serviceName + ".pid"
	return os.Remove(fileName)
}

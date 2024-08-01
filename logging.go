package godaemon

import (
	"log"
	"os"
)

func setupLogging(serviceName string) error {
	fileName := serviceName + ".log"

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(f)

	return nil
}

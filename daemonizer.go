package godaemon

import (
	"fmt"
	"os"
)

// Options for running the Daemonizer.
type Options struct {
	Name        string
	EnvVariable string
	Service     func() error
}

// Daemonizer is a structure to control the service. Use the New function to create it.
type Daemonizer struct {
	options Options
}

// New creates a Daemonizer.
func New(options Options) *Daemonizer {
	return &Daemonizer{
		options: options,
	}
}

// IsService should be executed at first in your main method.
// When the program is executed as the service, it executes your service function and then returns true.
// when the program is executed normally, it returns false, and you should use the Daemonizer to start or stop the
// service.
func (d *Daemonizer) IsService() (bool, error) {
	envVariable, err := validateEnvVariable(d.options.EnvVariable)
	if err != nil {
		return false, err
	}

	isDaemon := os.Getenv(envVariable) == "1"
	if isDaemon {
		serviceName := d.options.Name
		err := validateServiceName(serviceName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = writePidFile(serviceName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = setupLogging(serviceName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = d.options.Service()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return true, nil
	}

	return false, nil
}

// Start starts your program as background service.
func (d *Daemonizer) Start() error {
	envVariable, err := validateEnvVariable(d.options.EnvVariable)
	if err != nil {
		return err
	}

	return startProcess(envVariable)
}

// Stop tries to stop your running service.
func (d *Daemonizer) Stop() error {
	serviceName := d.options.Name
	err := validateServiceName(serviceName)
	if err != nil {
		return err
	}

	pid, err := readPidFile(serviceName)
	if err != nil {
		return err
	}

	err = stopProcess(pid)
	if err != nil {
		return err
	}

	return deletePidFile(d.options.Name)
}

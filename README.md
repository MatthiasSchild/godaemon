# godaemon

`godaemon` allows you to easily modify your application, like a web server, to be run in the background.
Specify your service, then implement the `Daemonizer`.
After that you only need to call the `Start` or `Stop` function depending on your conditions, like calling the cli
with specific parameters, and you're done.

It will start the program itself in a detached mode and pass an environment variable to separate the "controlling"
part from the "service" part.

## Daemonizer

The daemonizer is the tool to start and stop the daemon. You can call the `Start` or `Stop` methods to control the
service, e.g. using a CLI framework like Cobra, or just simply using `os.Args` like in the example.

    var daemonizer = godaemon.New(godaemon.Options{
        Name:        "service",
        EnvVariable: "DAEMON",
        Service:     service,
    })

- Name: The service name. This is required.
- EnvVariable: The name of the environment variable, to control if the program should run the service.
- Service: The function containing the service, which should run in the background.

## Service name

The service needs a name. The name will be part of the log and pid file.
The name should be build using letters, numbers and underscores, e.g. `web_server`.

## Log file

After the start, the `log` output is set to a log file. All log outputs will be written to the file.

## PID file

The service will write its process id into a file, so that the Daemonizer knows what process to kill when stopping.

## Example

In the example folder, a simple web server is implemented. It can be started or stopped using those keywords.

    // build the example
    go build -o example ./example

    // start the service
    ./example start
    // notice that the files "service.log" and "service.pid" are created

    // stop the service
    ./example stop

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/jeffbmartinez/cleanexit"
	"github.com/jeffbmartinez/delay"
	"github.com/jeffbmartinez/stdoutlog"

	"github.com/jeffbmartinez/todo-webserver/router"
)

const projectName string = "todo-webserver"
const defaultListenPort = 8000

func main() {
	cleanexit.SetUpSimpleExitOnCtrlC()

	allowAnyHostToConnect, listenPort := getCommandLineArgs()

	n := negroni.New()
	n.Use(delay.Middleware{})
	n.Use(stdoutlog.Middleware{})

	handlerRouter := router.GetRouter()
	n.UseHandler(handlerRouter)

	listenHost := "localhost"
	if allowAnyHostToConnect {
		listenHost = ""
	}
	displayServerInfo(listenHost, listenPort)

	listenAddress := fmt.Sprintf("%v:%v", listenHost, listenPort)
	n.Run(listenAddress)
}

func getCommandLineArgs() (allowAnyHostToConnect bool, port int) {
	flag.BoolVar(&allowAnyHostToConnect, "a", false, "Use to allow any ip address (any host) to connect. Default allows ony localhost.")
	flag.IntVar(&port, "port", defaultListenPort, "Port on which to listen for connections.")

	flag.Parse()

	/* Don't accept any positional command line arguments. flag.NArgs()
	counts only non-flag arguments. */
	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(2)
	}

	return
}

func displayServerInfo(listenHost string, listenPort int) {
	visibleTo := listenHost
	if visibleTo == "" {
		visibleTo = "All ip addresses"
	}

	fmt.Printf("%v is running.\n\n", projectName)
	fmt.Printf("Port: %v\n\n", listenPort)
	fmt.Printf("Hit [ctrl-c] to quit\n")
}

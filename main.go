package main

import (
	"github.com/urfave/cli"
	log "github.com/Sirupsen/logrus"
	"os"
	"fmt"
	"hans/mydocker/container"
)

const usage = "mydocker study. mydocker is a simple container runtime implementation"

var runCommand = cli.Command{
	Name:"run",
	Usage:`Create a container with namespace and cgroup limits
			mydocker run -it [command]
		`,
	Flags:[]cli.Flag{
		cli.BoolFlag{
			Name:"ti",
			Usage:"enable tty",
		},
	},

	Action:func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}

		cmd :=context.Args().Get(0)
		log.Infof("run command cmd=%s",cmd)
		tty :=context.Bool("ti")
		Run(tty,cmd)

		return nil
	},

}

var initCommand = cli.Command{
	Name:"init",
	Usage:"init container process run users process in container Do not call it outside",
	Action:func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("initCommand %s",cmd)
		err :=container.RunContainerInitProcess(cmd,nil)
		return err
	},
}

func main(){
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage


	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}


	app.Before = func(context *cli.Context) error {
		//Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(os.Stdout)
		return nil
	}

	log.Printf("os.Args=%s",os.Args)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}






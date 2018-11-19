package main

import (
	"flag"
	"io"
	"os"
	"os/exec"
)

var (
	ftl  = flag.Bool("ftl", false, "start the ftl server")
	eftl = flag.Bool("eftl", false, "start the eftl server")
)

func main() {
	flag.Parse()

	if *ftl {
		cmd := exec.Command("/opt/tibco/ftl/5.4/bin/tibrealmserver", "--http", "localhost:8080")
		stdout, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}
		err = cmd.Start()
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, stdout)
		err = cmd.Wait()
		if err != nil {
			panic(err)
		}
	} else if *eftl {
		cmd := exec.Command("/opt/tibco/eftl/3.4/ftl/bin/tibrealmadmin", "--realmserver", "http://localhost:8080",
			"--updaterealm", "/opt/tibco/eftl/3.4/samples/tibrealmserver.json")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}

		cmd = exec.Command("/opt/tibco/eftl/3.4/bin/tibeftlserver", "--realmserver", "http://localhost:8080",
			"--listen", "ws://localhost:9191")
		stdout, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}
		err = cmd.Start()
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, stdout)
		err = cmd.Wait()
		if err != nil {
			panic(err)
		}
	} else {
		flag.PrintDefaults()
	}
}

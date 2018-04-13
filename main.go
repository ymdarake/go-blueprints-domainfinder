package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{
	exec.Command("lib/synonyms"),
	exec.Command("lib/sprinkle"),
	exec.Command("lib/coolify"),
	exec.Command("lib/domainify"),
	exec.Command("lib/whois"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout

	// setup  pipe chain
	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		pipedStdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Panicln(err)
		}
		nextCmd.Stdin = pipedStdout
	}

	// run each in background
	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Panicln(err)
		} else {
			defer cmd.Process.Kill()
		}
	}

	// wait for each to be done (before exiting from this main process)
	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Panicln(err)
		}
	}

}

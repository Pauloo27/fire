package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var (
	commandToRun string
	rootFolder   string
	ignorePaths  string
)

func init() {
	flag.StringVar(&commandToRun, "command", "", "The command to run")
	flag.StringVar(&rootFolder, "root", ".", "The root folder to watch")
	flag.StringVar(&ignorePaths, "ignore", "", "The paths to ignore, split by comma")

	flag.Parse()
	if commandToRun == "" {
		flag.Usage()
		return
	}
}

var (
	cmd *exec.Cmd
)

func callCommand() {
	if cmd != nil {
		_ = cmd.Process.Kill()
	}
	splitted := strings.Split(commandToRun, " ")
	cmd = exec.Command(splitted[0], splitted[1:]...)
	_ = cmd.Run()
}

func main() {
	if s, err := os.Stat(rootFolder); os.IsNotExist(err) || !s.IsDir() {
		fmt.Printf("The root folder %s does not exist or is not a folder\n", rootFolder)
		return
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	// TODO: do it properly
	err = watcher.Add(rootFolder)
	if err != nil {
		panic(err)
	}
	for {
		event, ok := <-watcher.Events
		if !ok {
			panic("watcher error")
		}
		// TODO: allow custom events
		if event.Op == fsnotify.Chmod {
			continue
		}
		// TODO: debounce
		callCommand()
	}
}

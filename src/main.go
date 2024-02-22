package main

import (
	"imgedit/src/cli"
	"imgedit/src/web"
	"os"
)

func main() {
    args := os.Args[1:]
    if len(args) > 0 && args[0] == "cli" {
        cli.RunCli()
        return
    }
    web.RunWeb()
}


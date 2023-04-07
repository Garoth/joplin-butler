package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Garoth/joplin-butler/endpoints"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmsgprefix)
	log.SetPrefix("> ")

	authSet := flag.NewFlagSet("auth", flag.ExitOnError)

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "auth":
		authSet.Parse(os.Args[2:])
		if err := endpoints.Auth(); err != nil {
			log.Println("ERR:", err)
		}

	default:
		printHelp()
		os.Exit(1)
	}

	os.Exit(0)
}

func printHelp() {
	fmt.Println("Available Subcommands:\n" +
		"   - auth\n" +
		"   - help")
}

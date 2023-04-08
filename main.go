package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Garoth/joplin-butler/endpoints"
	"github.com/Garoth/joplin-butler/types"
	"github.com/Garoth/joplin-butler/utils"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmsgprefix)
	log.SetPrefix("> ")
	// TODO:lung:2023-04-07 create a flag to disable debug output
	log.SetOutput(ioutil.Discard)

	authSet := flag.NewFlagSet("auth", flag.ExitOnError)

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	if err := endpoints.Auth(); err != nil {
		log.Fatalln("ERR:", err)
	}

	switch os.Args[1] {
	case "auth":
		authSet.Parse(os.Args[2:])
		os.Exit(1)
	case "notes":
		notesStr, err := utils.GetPath("notes")
		if err != nil {
			log.Fatalln("ERR:", err)
		}
		notes, err := types.NewPaginated[types.Note](notesStr)
		if err != nil {
			log.Fatalln("ERR:", err)
		}
		fmt.Println(notes)
	case "note":
		notesStr, err := utils.GetPath("notes/1b16f5bfd4854f1f98345ef3d5dcc8c3" +
			"?fields=id,title,parent_id,created_time,updated_time,source,body")
		if err != nil {
			log.Fatalln("ERR:", err)
		}
		var note types.Note
		err = json.Unmarshal([]byte(notesStr), &note)
		if err != nil {
			log.Fatalln("ERR:", err)
		}
		fmt.Println(note.DetailedString())

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

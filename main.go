package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Garoth/joplin-butler/endpoints"
	"github.com/Garoth/joplin-butler/types"
	"github.com/Garoth/joplin-butler/utils"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmsgprefix)
	log.SetPrefix("> ")
	// TODO:lung:2023-04-07 create a flag to disable debug output
	// log.SetOutput(ioutil.Discard)

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

	case "get":
		if len(os.Args) < 3 {
			printHelp()
			os.Exit(1)
		}

		itemTypeStr := os.Args[2]
		itemId := ""
		if strings.Contains(itemTypeStr, "/") {
			parts := strings.Split(itemTypeStr, "/")
			itemTypeStr = parts[0]
			itemId = parts[1]
		} else if len(os.Args) >= 4 {
			itemId = os.Args[3]
		}

		itemTypeID, err := types.NewItemTypeID(itemTypeStr)
		if err != nil {
			// try again without the last character which might be an 's'
			var err2 error
			itemTypeID, err2 = types.NewItemTypeID(itemTypeStr[0 : len(itemTypeStr)-1])
			if err2 != nil {
				log.Fatalln("ERR:", err)
			}
		}

		// Query string is handled differently for non-note types, and
		// we need to use the wildcard to search for all
		if itemTypeID != types.ItemTypeNote && itemId == "" {
			itemId = "*"
		}
		res, err := utils.GetPath("search?query=" + itemId +
			"&type=" + itemTypeID.String())
		if err != nil {
			log.Fatalln("ERR:", err)
		}
		jsonErr, err := types.NewError(res)
		if err == nil {
			log.Fatalln("ERR:", jsonErr)
		}

		log.Println("type", itemId, itemTypeID)

		items, err := types.NewPaginated[types.ItemInfo](res)
		if err != nil {
			log.Fatalln("ERR:", err)
		}
		fmt.Println(items)

	case "note":
		if len(os.Args) < 3 {
			log.Fatalln("ERR: must have more parameters for note lookup")
		}
		targetNote := os.Args[2]
		notesStr, err := utils.GetPath("notes/" + targetNote +
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

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
	createSet := flag.NewFlagSet("create", flag.ExitOnError)
	createBody := createSet.String("body", "", "For notes, the markdown body")
	createHTMLBody := createSet.String("body_html", "", "For notes, the html body")

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
		os.Exit(0)

	case "get":
		itemTypeID, itemName, _, err := parseItemTypeAndName(os.Args)
		if err != nil  {
			log.Fatalln("ERR:", err)
		}

		if itemTypeID == types.ItemTypeNote  && itemName == "" {
			items, err := listNotes()
			if err != nil {
				log.Fatalln("ERR:", err)
			}
			fmt.Println(items)
		} else if itemTypeID == types.ItemTypeNote  && itemName != "" {
			notesStr, err := utils.GetPath("notes/" + itemName +
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
		} else {
			items, err := queryItemType(itemTypeID, itemName)
			if err != nil {
				log.Fatalln("ERR:", err)
			}
			fmt.Println(items)
		}

	case "create":
		itemTypeID, itemName, remArgs, err := parseItemTypeAndName(os.Args)
		if err != nil  {
			log.Fatalln("ERR:", err)
		}
		createSet.Parse(remArgs)

		switch itemTypeID {
		case types.ItemTypeNote:
			dataJson := fmt.Sprintf("{ \"title\": \"%s\"", itemName)
			if *createBody != "" {
				dataJson += ", \"body\": \"" + *createBody + "\""
			}
			if *createHTMLBody != "" {
				dataJson += ", \"body_html\": \"" + *createHTMLBody + "\""
			}
			dataJson += " }"

			res, err := utils.PostPath("notes", dataJson)
			err = types.CheckError(res, err)
			if err != nil {
				log.Fatalln("ERR:", err)
			}
			note := &types.Note{}
			err = json.Unmarshal([]byte(res), note)
			if err != nil  {
				log.Fatalln("ERR:", err)
			}
			fmt.Println(note.DetailedString())

		default:
			endpoint := itemTypeID.String() + "s"
			res, err := utils.PostPath(endpoint,
				fmt.Sprintf("{ \"title\": \"%s\" }", itemName))
			err = types.CheckError(res, err)
			if err != nil {
				log.Fatalln("ERR:", err)
			}
			item := &types.ItemInfo{}
			err = json.Unmarshal([]byte(res), item)
			if err != nil  {
				log.Fatalln("ERR:", err)
			}
			log.Printf("Successfully created '%s/%s'", endpoint, itemName)
			fmt.Println(item.String())
		}

	case "delete":
		itemTypeID, itemName, _, err := parseItemTypeAndName(os.Args)
		if err != nil  {
			log.Fatalln("ERR:", err)
		}
		if itemName == "" {
			log.Fatalln("ERR: must specify item id to delete")
		}
		endpoint := itemTypeID.String() + "s"
		res, err := utils.DeletePath(endpoint + "/" + itemName)
		err = types.CheckError(res, err)
		if err != nil {
			log.Fatalln("ERR:", err)
		}
		log.Println("Successfully deleted " + endpoint + "/" + itemName)

	default:
		printHelp()
		os.Exit(1)
	}

	os.Exit(0)
}

func printHelp() {
	fmt.Println("Available Subcommands Examples:\n" +
		"   - auth\n" +
		"   - get notes/eb5e1e29c7164c4d9b9ed4a11d218cdc\n" +
		"   - create 'note/my title here' -body '# Cool heading\n\nMore'\n" +
		"   - delete note/eb5e1e29c7164c4d9b9ed4a11d218cdc\n" +
		"   - edit (TODO)\n" +
		"   - attach (TODO)\n" +
		"   - remove (TODO)\n" +
		"   - help")
}

// Parses an item spec like "tag/address" or "tags" or "tag address"
func parseItemTypeAndName(args []string) (types.ItemTypeID, string, []string, error) {
	if len(args) < 3 {
		return types.ItemTypeNone, "", []string{}, fmt.Errorf("not enough arguments")
	}

	lastArg := 2
	itemTypeStr := args[2]
	itemName := ""
	if strings.Contains(itemTypeStr, "/") {
		parts := strings.Split(itemTypeStr, "/")
		itemTypeStr = parts[0]
		itemName = parts[1]
	} else if len(args) >= 4 {
		itemName = args[3]
		lastArg = 3
	}

	itemTypeID, err := types.NewItemTypeID(itemTypeStr)
	if err != nil {
		// try again without the last character which might be an 's'
		var err2 error
		itemTypeID, err2 = types.NewItemTypeID(itemTypeStr[0 : len(itemTypeStr)-1])
		if err2 != nil {
			return types.ItemTypeNone, "", []string{}, err
		}
	}

	remainingArgs := []string{}
	if len(args) > lastArg+1 {
		remainingArgs = args[lastArg+1:len(args)]
	}

	return itemTypeID, itemName, remainingArgs, nil
}

func listNotes() (*types.Paginated[types.Note], error) {
	allItems := &types.Paginated[types.Note]{
		Items: []types.Note{},
		HasMore: false,
	}

	for page := 1; ; page++ {
		res, err := utils.GetPath(utils.AppendQueryStringToPath("notes", "page", page))
		err = types.CheckError(res, err)
		if err != nil {
			return nil, err
		}

		items, err := types.NewPaginated[types.Note](res)
		if err != nil {
			return nil, err
		}

		allItems.Items = append(allItems.Items, items.Items...)

		if !items.HasMore {
			break
		}
	}

	return allItems, nil
}

// For non-note item types, we can query them by title using search
func queryItemType(itemTypeID types.ItemTypeID,
	itemName string) (*types.Paginated[types.ItemInfo], error) {

	if itemName == "" {
		itemName = "*"
	}
	allItems := &types.Paginated[types.ItemInfo]{
		Items: []types.ItemInfo{},
		HasMore: false,
	}

	for page := 1; ; page++ {
		query := utils.AppendQueryStringsToPath("search", map[string]any{
			"query": itemName,
			"type": itemTypeID.String(),
			"page": page,
		})
		log.Println("Query:", query)
		res, err := utils.GetPath(query)
		err = types.CheckError(res, err)
		if err != nil {
			return nil, err
		}

		items, err := types.NewPaginated[types.ItemInfo](res)
		if err != nil {
			return nil, err
		}

		allItems.Items = append(allItems.Items, items.Items...)

		if !items.HasMore {
			break
		}
	}

	return allItems, nil
}

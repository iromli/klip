package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	tm "github.com/buger/goterm"
	"github.com/docopt/docopt-go"
	"github.com/spf13/afero"
)

func interfaceToString(val interface{}) string {
	switch val.(type) {
	case string:
		return val.(string)
	case []string:
		return strings.Join(val.([]string), " ")
	default:
		return ""
	}
}

func getFilepath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	filepath := path.Join(u.HomeDir, ".klip")
	return filepath, nil
}

func printErr(err error) {
	fmt.Println(tm.Color(fmt.Sprintf("[ERR] %s", err), tm.RED))
}

func main() {
	usage := `Klip

Usage:
  klip put <list> [(<name> <value>...)]
  klip get <list> [<name>]
  klip delete <list> [<name>]

`
	args, _ := docopt.Parse(usage, nil, true, "Klip 0.1", false)

	list := interfaceToString(args["<list>"])
	name := interfaceToString(args["<name>"])
	value := interfaceToString(args["<value>"])

	fp, err := getFilepath()
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	s, err := NewJSONStorage(fp, afero.OsFs{})
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	switch {
	case args["put"]:
		if err := s.Put(list, name, value); err != nil {
			printErr(err)
			os.Exit(1)
		}
	case args["get"].(bool) && name != "":
		result, err := s.Get(list, name)
		if err != nil {
			printErr(err)
			os.Exit(1)
		}
		fmt.Println(result)
	case args["get"].(bool) && name == "":
		result, err := s.Map(list)
		if err != nil {
			printErr(err)
			os.Exit(1)
		}
		fmt.Println(tm.Color(fmt.Sprintf("%s (%d)", list, len(result)), tm.GREEN))
		for k, v := range result {
			fmt.Printf("  %s: %s\n", tm.Bold(k), v)
		}
	case args["delete"]:
		if err := s.Delete(list, name); err != nil {
			printErr(err)
			os.Exit(1)
		}
	}
}

package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/iromli/clip/storage"
)

func interfaceToString(val interface{}) string {
	if val != nil {
		return val.(string)
	}
	return ""
}

func main() {
	usage := `Clip

Usage:
  clip put <list> [(<name> <value>)]
  clip get <list> [<name>]
  clip delete <list> [<name>]

`
	args, _ := docopt.Parse(usage, nil, true, "Clip 0.1", false)
	list := interfaceToString(args["<list>"])
	name := interfaceToString(args["<name>"])
	value := interfaceToString(args["<value>"])

	s, err := storage.NewJSONStorage()
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case args["put"]:
		if err := s.Put(list, name, value); err != nil {
			log.Fatal(err)
		}
	case args["get"]:
		if name != "" {
			result, err := s.Get(list, name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
		} else {
			result, err := s.List(list)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
		}
	case args["delete"]:
		if err := s.Delete(list, name); err != nil {
			log.Fatal(err)
		}
	}
}

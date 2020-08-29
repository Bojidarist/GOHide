package main

import (
	"gohide/algos"
	"fmt"
	"os"
)

func showHelp() {
	fmt.Println("Options:\n-se [img] [output img] [ASCII string] (Simple Encode)\n-sd [img] (Simple Decode)\n")
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-se" {
			if len(os.Args) > 4 {
				algos.SimpleAlgoEncode(os.Args[2], os.Args[3], os.Args[4])
			} else { showHelp() }
		} else if os.Args[1] == "-sd" {
			if len(os.Args) > 2 {
				fmt.Println(algos.SimpleAlgoDecode(os.Args[2]))
			} else { showHelp() }
		}
	} else { showHelp() }
}

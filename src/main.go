package main

import (
	"strconv"
	"gohide/algos"
	"fmt"
	"os"
)

func showHelp() {
	fmt.Println("Options:\n-se [img] [output img] [ASCII string] (Simple Encode)\n-sd [img] [skip] (Simple Decode)\n")
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-se" {
			if len(os.Args) > 4 {
				algos.SimpleAlgoEncode(os.Args[2], os.Args[3], os.Args[4])
			} else { showHelp() }
		} else if os.Args[1] == "-sd" {
			if len(os.Args) > 3 {
				skip, _ := strconv.Atoi(os.Args[3])
				fmt.Println(algos.SimpleAlgoDecode(os.Args[2], skip))
			} else { showHelp() }
		}
	} else { showHelp() }
}

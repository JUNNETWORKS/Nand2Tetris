package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	filePath := flag.Arg(0)
	parser, _ := NewParser(filePath)

	fmt.Printf("&Parser: \n\t%v\n", parser)
	fmt.Printf("&Parser.Commands: \n\t%v\n", parser.Commands)
	for i, command := range parser.Commands {
		fmt.Printf("LINE%d:\t%#v\n", i, command)
	}
	code := NewCode()
	for i, command := range parser.Commands {
		binaryInstruction, err := code.Assemble(command)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d:\t%s\n", i, binaryInstruction)
	}
}

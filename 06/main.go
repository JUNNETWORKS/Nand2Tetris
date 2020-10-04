package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	err := os.Remove("output.hack")
	hackFile, err := os.OpenFile("output.hack", os.O_WRONLY|os.O_CREATE, 0666)
	defer hackFile.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	code := NewCode()
	for i, command := range parser.Commands {
		binaryInstruction, err := code.Assemble(command)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d:\t%s\n", i, binaryInstruction)
		_, err = hackFile.WriteString(fmt.Sprintf("%s\n", binaryInstruction))
		if err != nil {
			log.Fatal(err)
		}
	}
}

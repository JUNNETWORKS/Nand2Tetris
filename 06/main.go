package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var symbolTable *SymbolTable

func init() {
	symbolTable = NewSymbolTable()
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func main() {
	flag.Parse()
	filePath := flag.Arg(0)

	parser, _ := NewParser(filePath)
	fmt.Printf("&Parser: \n\t%v\n", parser)
	fmt.Printf("&Parser.Commands: \n\t%v\n", parser.Commands)
	for i, command := range parser.Commands {
		fmt.Printf("LINE%d:\t%#v\n", i, command)
	}

	hackFilePath := fmt.Sprintf("%s.hack", fileNameWithoutExtension(filepath.Base(filePath)))
	err := os.Remove(hackFilePath)
	hackFile, err := os.OpenFile(hackFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer hackFile.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%#v\n", symbolTable)

	code := NewCode()
	for i, command := range parser.Commands {
		if command.Type == L_COMMAND {
			continue
		}
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// parser : Parserの主な機能は各アセンブリコマンドをその基本要素(フィールドとシンボル)に分解することである.
type parser interface {
	HasMoreCommands() bool
	Advance()
	CommandType() int
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}

// Parse : 行を解析して, Commandオブジェクトのポインタを返す
func Parse(line string) *Command {
	// Remove spaces
	line = strings.ReplaceAll(line, " ", "")

	// 空行だったらすぐ返す
	if len(line) == 0 {
		return nil
	}

	// Comment
	if line[0:2] == "//" {
		return nil
	}

	// コメントは事前に消しておく
	commentIdx := strings.Index(line, "//")
	if commentIdx > 0 {
		line = line[:commentIdx]
	}

	instruction := new(Command)

	// A命令
	if line[0] == '@' {
		instruction.Type = A_COMMAND
		instruction.Value = line[1:]
		return instruction
	}

	// C命令 or ラベル
	var tmp []byte
	nextArea := ""

	for _, r := range line {
		if r == '=' {
			instruction.Dest = string(tmp)
			tmp = tmp[:0] // tmp を空に (capはそのまま)
			nextArea = "comp"
		} else if r == ';' {
			instruction.Comp = string(tmp)
			tmp = tmp[:0] // tmp を空に (capはそのまま)
			if len(instruction.Dest) == 0 {
				instruction.Dest = "null"
			}
			nextArea = "jump"
		} else {
			tmp = append(tmp, byte(r))
		}
	}

	if nextArea == "comp" {
		instruction.Comp = string(tmp)
		instruction.Jump = "null"
		instruction.Type = C_COMMAND
	} else if nextArea == "jump" {
		instruction.Jump = string(tmp)
		instruction.Type = C_COMMAND
	} else {
		// TODO: ラベル対応
		instruction.Type = L_COMMAND
	}
	return instruction
}

// Parser Parser構造体.
// Parser構造体はアセンブリファイルを解析し, それをコマンドオブジェクトの配列として保管する.
// Indexプロパティは現在対象としているコマンドのインデックスを表す.
type Parser struct {
	Index    int
	Commands []*Command
}

// ParserError Parser関連のエラーを表す
type ParserError struct {
	Message string
}

func (e *ParserError) Error() string {
	return e.Message
}

// NewParser 新しいParserオブジェクトのポインタを返す
func NewParser(filePath string) (*Parser, error) {
	if filepath.Ext(filePath) != ".asm" {
		err := ParserError{Message: fmt.Sprintf("%s is not a assembly file.", filePath)}
		return nil, &err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create ROM symbol table
	fileScanner := bufio.NewScanner(file)
	// currentRomAddress := 0
	// for fileScanner.Scan() {
	// 	line := fileScanner.Text()
	// 	line = strings.TrimSpace(line)
	// 	if line[0] == '(' {
	// 		symbol := strings.TrimSuffix(line, ")")
	// 		symbolTable.AddEntry(symbol, currentRomAddress)
	// 	} else {
	// 		currentRomAddress++
	// 	}
	// }

	// parse commands
	file.Seek(0, 0)
	parser := new(Parser)
	// currentRamAddress := 16
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		instruction := Parse(line)
		if instruction != nil {
			parser.Commands = append(parser.Commands, instruction)
		}
	}
	return parser, nil
}

// HasMoreCommands : まだコマンドが残っているか
func (p *Parser) HasMoreCommands() bool {
	return len(p.Commands) > p.Index
}

// Advance : 次のコマンドに対象を移す
func (p *Parser) Advance() {
	p.Index++
}

// CommandType : 現在対象としているコマンドのコマンドタイプを取得する
func (p *Parser) CommandType() int {
	return p.Commands[p.Index].Type
}

// Symbol : A命令, もしくはラベル, 変数 のシンボルを取得.
// これは10進数か文字列をstring型で返す.
// A_COMMAND か L_COMMAND のときだけ呼ばれる
func (p *Parser) Symbol() string {
	return p.Commands[p.Index].Value
}

// Dest : C命令のDestの値を取得する.
// C_COMMAND のときだけ呼ばれる
func (p *Parser) Dest() string {
	return p.Commands[p.Index].Dest
}

// Comp : C命令のCompの値を取得する.
// C_COMMAND のときだけ呼ばれる
func (p *Parser) Comp() string {
	return p.Commands[p.Index].Comp
}

// Jump : C命令のJumpの値を取得する.
// C_COMMAND のときだけ呼ばれる
func (p *Parser) Jump() string {
	return p.Commands[p.Index].Jump
}

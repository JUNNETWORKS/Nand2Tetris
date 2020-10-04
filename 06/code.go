package main

import (
	"fmt"
	"strconv"
)

// Code : Command構造体を元にバイナリコードを生成する構造体
// Hackのアセンブリ言語のニーモニックをバイナリコードへ変換する
type Code struct {
	DestTable map[string]string
	CompTable map[string]string
	JumpTable map[string]string
}

// NewCode : 各バイナリ命令テーブル登録済みのCode構造体のポインタを返す
func NewCode() *Code {
	code := new(Code)
	code.DestTable = map[string]string{
		"null": "000",
		"M":    "001",
		"D":    "010",
		"MD":   "011",
		"A":    "100",
		"AM":   "101",
		"AD":   "110",
		"AMD":  "111",
	}
	code.CompTable = map[string]string{
		"0":   "0101010",
		"1":   "0111111",
		"-1":  "0111010",
		"D":   "0001100",
		"A":   "0110000",
		"!D":  "0001101",
		"!A":  "0110001",
		"-D":  "0001111",
		"-A":  "0110011",
		"D+1": "0011111",
		"A+1": "0110111",
		"D-1": "0001110",
		"A-1": "0110010",
		"D+A": "0000010",
		"D-A": "0010011",
		"A-D": "0000111",
		"D&A": "0000000",
		"D|A": "0010101",
		"M":   "1110000",
		"!M":  "1110001",
		"-M":  "1110011",
		"M+1": "1110111",
		"M-1": "1110010",
		"D+M": "1000010",
		"D-M": "1010011",
		"M-D": "1000111",
		"D&M": "1000000",
		"D|M": "1010101",
	}
	code.JumpTable = map[string]string{
		"null": "000",
		"JGT":  "001",
		"JEQ":  "010",
		"JGE":  "011",
		"JLT":  "100",
		"JNE":  "101",
		"JLE":  "110",
		"JMP":  "111",
	}
	return code
}

// CodeError : Code から排出されるエラー
type CodeError struct {
	Message string
}

func (e *CodeError) Error() string {
	return e.Message
}

// Dest : 引数の文字列から対応するバイナリを取得
// 引数はニーモニックを表す文字列
// 返り値のstringはビットを表す文字列  ex: "001101"
func (code *Code) Dest(mnemonic string) (string, error) {
	bin, ok := code.DestTable[mnemonic]
	if !ok {
		return "", &CodeError{Message: fmt.Sprintf("%s is not in DestTable", mnemonic)}
	}
	return bin, nil
}

// Comp : 引数の文字列から対応するバイナリを取得
func (code *Code) Comp(mnemonic string) (string, error) {
	bin, ok := code.CompTable[mnemonic]
	if !ok {
		return "", &CodeError{Message: fmt.Sprintf("%s is not in CompTable", mnemonic)}
	}
	return bin, nil
}

// Jump : 引数の文字列から対応するバイナリを取得
func (code *Code) Jump(mnemonic string) (string, error) {
	bin, ok := code.JumpTable[mnemonic]
	if !ok {
		return "", &CodeError{Message: fmt.Sprintf("%s is not in JumpTable", mnemonic)}
	}
	return bin, nil
}

// Assemble : Commandオブジェクトからバイナリ命令を表す文字列を作成して返す
func (code *Code) Assemble(command *Command) (string, error) {
	binInstruction := ""

	if command.Type == A_COMMAND {
		binInstruction += "0"
		i, err := strconv.ParseInt(command.Value, 10, 64) // 文字列からint64に変換
		if err != nil {
			return "", err
		}
		bin := strconv.FormatInt(i, 2)
		// 15桁0埋め
		for i := 0; i < 15-len(bin); i++ {
			binInstruction += "0"
		}
		binInstruction += bin

		// TODO: ラベル対応

	} else if command.Type == C_COMMAND {
		binInstruction += "111"
		if bin, err := code.Comp(command.Comp); err == nil {
			binInstruction += bin
		} else {
			return "", err
		}
		if bin, err := code.Dest(command.Dest); err == nil {
			binInstruction += bin
		} else {
			return "", err
		}
		if bin, err := code.Jump(command.Jump); err == nil {
			binInstruction += bin
		} else {
			return "", err
		}
	}
	return binInstruction, nil
}

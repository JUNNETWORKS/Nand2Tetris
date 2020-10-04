package main

const (
	A_COMMAND = iota // A命令
	C_COMMAND        // C命令
	L_COMMAND        // 疑似コマンド
)

// Command : 各行の命令を表す構造体
type Command struct {
	// Is the instruction C or A?
	Type int

	// A instruction
	Value string

	// C instruction
	Dest string
	Comp string
	Jump string
}

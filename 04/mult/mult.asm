// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// 0 <= R0 <= R1

// Put your code here.
@sum  // Aレジスタに sum という変数の値を入れる. sum変数の値(アドレス)は自動的にユニークなものがアセンブラから付けられる
M = 0  // M = Memory[A]

// ループの回数を記録
@R1
D = M
@loop_count 
M = D

(LOOP)
    @loop_count 
    D = M
    @LOOP
    D;JGT  // loop_count が0より上ならLOOPに戻る
    
    @R0
    D = M
    @sum
    M = M + D  // sum = sum + R0
    
    @loop_count
    M = M - 1

// 結果をR2に入れる
@sum
D = M
@R2
M = D
(END)
    @END
    0;JMP  // 無限ループ
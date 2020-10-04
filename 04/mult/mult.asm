// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
// 0 <= R0 <= R1

// Put your code here.

// R2に合計値を入れるので0で初期化
@R2
M = 0

(LOOP)
    // ループの条件判定
    @R1
    D = M
    @END
    D;JLE  // R1の値が0以下ならループ終了

    @R0
    D = M
    @R2
    M = D + M  // R2 = R0 + R2
    
    @R1
    M = M - 1
    
    @LOOP
    0;JMP
(END)
    @END
    0;JMP  // 無限ループ
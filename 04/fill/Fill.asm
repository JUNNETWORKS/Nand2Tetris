// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

@8192  // 横32レジスタ * 縦256行 分埋める
D = A
@end_of_screen_word
M = D

(LOOP)
    // キーボードの入力を取得
    @KBD  // キーボードの I/Oポインタ のRAMアドレスはシンボルとして定義されている 24576(0x6000)
    D = M

    @FILL_BLACK
    D;JNE  // 0じゃないならジャンプ
    @FILL_WHITE
    D;JEQ  // 0ならジャンプ
    
    (END_FILL)

    @LOOP
    0;JMP

// 画面を全て黒(1)で塗りつぶす
(FILL_BLACK)
    // スクリーンI/Oポインタ は 16384(0x4000)
    // スクリーンは 横512 * 縦256 ピクセルの白黒
    // 512 = 16 * 32,    256 = 16 * 16
    // r番目の行のc番目の列のピクセルは RAM[16384 + r * 32 + c / 16] のc%16番目のビット
    // スクリーンの I/Oポインタ はシンボルとして定義されている 16384(0x4000)

    @current_idx
    M = 0    // ピクセルを指定する用
    (FILL_BLACK_LOOP)
        @current_idx
        D = M
        @SCREEN
        A = A + D        // SCREEN + idx で現在のメモリのワードのアドレス取得
        M = -1             // M = 0xffff  // 黒塗り
        @current_idx
        MD = M + 1        // current_position = current_position + 1
        
        @end_of_screen_word
        D = M - D

        @FILL_BLACK_LOOP
        D;JGE

    @END_FILL
    0;JMP

// 画面を全て白(0)で塗りつぶす
(FILL_WHITE)
    @current_idx
    M = 0    // ピクセルを指定する用
    (FILL_WHITE_LOOP)
        @current_idx
        D = M
        @SCREEN
        A = A + D        // SCREEN + idx で現在のメモリのワードのアドレス取得
        M = 0             // M = 0x0000  // 白塗り
        @current_idx
        MD = M + 1        // current_position = current_position + 1
        
        @end_of_screen_word
        D = M - D

        @FILL_WHITE_LOOP
        D;JGE

    @END_FILL
    0;JMP
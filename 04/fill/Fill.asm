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

D = 65535  // 0xffff
@black_16
M = D

D = 0      // 0x0000
@white_16
M = D

D = 16896  // 0x4200
@end_of_screen
M = D

(LOOP)
    // キーボードの入力を取得
    @KBD  // キーボードの I/Oポインタ のRAMアドレスはシンボルとして定義されている 24576(0x6000)
    D = M

    @FILL
    D;JNE  // 0じゃないならジャンプ
    @FILL_WHITE
    D;JEQ  // 0ならジャンプ
    
    (END_FILL)

    @LOOP
    0;JMP

// 画面を全て0か1で塗りつぶす
(FILL)
    // スクリーンI/Oポインタ は 16384(0x4000)
    // スクリーンは 横512 * 縦256 ピクセルの白黒
    // 512 = 16 * 32,    256 = 16 * 16
    // r番目の行のc番目の列のピクセルは RAM[16384 + r * 32 + c / 16] のc%16番目のビット

    @SCREEN  // スクリーンの I/Oポインタ はシンボルとして定義されている 16384(0x4000)
    D = A
    @screen_idx
    M = 0    // SCREEN + idx でピクセルを指定する用
    (FILL_LOOP)
        @current_position
        D = M
        @black_16
        D = M            // D = 0xffff
        @current_position
        M = D            // Memory[*(current_position)] = 0xffff  // ex(初回): Memory[0x4000] = 0xffff
        @current_position
        M = A + 1        // current_position = current_position + 1
        
        // 0x4200までループ
        @end_of_screen
        D = M
        @current_position
        D = D - M        // 0x4200 - current_position
        @FILL_LOOP
        D;JGE

    @END_FILL
    0:JMP

// 画面を全て白(0)で塗りつぶす
(FILL_WHITE)
    @current_position
    M = SCREEN
    (FILL_WHITE_LOOP)
        @current_position
        D = M
        @D
        M = 0  // 0x0000
        @current_position
        M = M + 1  // current_position += 1
        
        // 0x4200までループ
        D = 16896 - M 
        @FILL_LOOP
        D;JGE

    @END_FILL
    0:JMP
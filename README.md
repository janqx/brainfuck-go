# Brainfuck interpreter in Go

## Feature
- REPL
- Executable String
- Runable File
- Context
- Easy to read the code

## Quick Start
```
go get github.com/janqx/brainfuck-go
```

## Simple Example
```bash
cd brainfuck-go

# REPL
go run cli/main
>> ++++++++++++++++++++++++++++++++++++
>> .
$  
>> +.
%  
>> !exit

# Execute String
go run cli/main.go -c "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
Output: Hello World!

# Run File
go run cli/main.go example/mandelbrot_tiny.bf
Output:
AAAAAAAABBBBBBBBCCCCCCCCCCCCCCCCCCDDDDEFEEDDDCCCCCBBBBBBBBBBBBBBB
AAAAAAABBBBBBCCCCCCCCCCCCCCCCCDDDDDDEEFIKGGGDDDDDCCCCBBBBBBBBBBBB
AAAAAABBBBCCCCCCCCCCCCCCCCCDDDDDDDEEEFGHKPIGFEDDDDDCCCCCBBBBBBBBB
AAAAABBBCCCCCCCCCCCCCCCCCDDDDDDDEEEFGPVT  Q[HEEEEDDDCCCCCCBBBBBBB
AAAABBCCCCCCCCCCCCCCCCDDDDDDDEEFFFGGHK      HGFFEEEDDDCCCCCBBBBBB
AAABBCCCCCCCCCCCCCCCDDDDDEEEFGK MJJ NR    YS L HHGIJFDDCCCCCCBBBB
AAABCCCCCCCCCCCCCDDDEEEEEEFFFHI                    MGEDDCCCCCCBBB
AABCCCCCCCCCCCDDEEEEEEEEFFFGY Q                   MHGEEDCCCCCCCBB
AACCCCCCDDDDDEEFLHGGHMHGGGHIR                      QLHEDDCCCCCCCB
ABCCDDDDDDEEEEFGIKU    RLJJL                        IFEDDCCCCCCCB
ACDDDDDDEEEEEGGHOS        QR                        JFEDDDCCCCCCC
ADDDDDEFFFGGHKOPS                                   GEEDDDCCCCCCC
A                                                PJGFEEDDDCCCCCCC
ADDDDDEFFFGGHKOPS                                   GEEDDDCCCCCCC
ACDDDDDDEEEEEGGHOS        QR                        JFEDDDCCCCCCC
ABCCDDDDDDEEEEFGIKU    RLJJL                        IFEDDCCCCCCCB
AACCCCCCDDDDDEEFLHGGHMHGGGHIR                      QLHEDDCCCCCCCB
AABCCCCCCCCCCCDDEEEEEEEEFFFGY Q                   MHGEEDCCCCCCCBB
AAABCCCCCCCCCCCCCDDDEEEEEEFFFHI                    MGEDDCCCCCCBBB
AAABBCCCCCCCCCCCCCCCDDDDDEEEFGK MJJ NR    YS L HHGIJFDDCCCCCCBBBB
AAAABBCCCCCCCCCCCCCCCCDDDDDDDEEFFFGGHK      HGFFEEEDDDCCCCCBBBBBB
AAAAABBBCCCCCCCCCCCCCCCCCDDDDDDDEEEFGPVT  Q[HEEEEDDDCCCCCCBBBBBBB
AAAAAABBBBCCCCCCCCCCCCCCCCCDDDDDDDEEEFGHKPIGFEDDDDDCCCCCBBBBBBBBB
AAAAAAABBBBBBCCCCCCCCCCCCCCCCCDDDDDDEEFIKGGGDDDDDCCCCBBBBBBBBBBBB

```

## References
- https://brainfuck.org/
- https://github.com/fabianishere/brainfuck